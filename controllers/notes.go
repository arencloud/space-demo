package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/arencloud/space-demo/initlib"
	"github.com/arencloud/space-demo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	var payload *models.CreateNoteSchema

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()

	note := models.Note {
		Title: payload.Title,
		Content: payload.Content,
		Category: payload.Category,
		Published: payload.Published,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res := initlib.DB.Create(&note)

	if res.Error != nil && strings.Contains(res.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exists"})
	} else if res.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": note}})
}

func SearchNotesHandler(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	var notes []models.Note

	initPage, _ := strconv.Atoi(page)
	initLimit, _ := strconv.Atoi(limit)
	offset := (initPage - 1) * initLimit

	res := initlib.DB.Limit(initLimit).Offset(offset).Find(&notes)
	if res.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": res.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(notes), "notes": notes})
}

func UpdateNoteHandler(c *fiber.Ctx) error {
	nodeId := c.Params("noteId")
	var note models.Note
	var payload *models.UpdateNoteSchema

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	res := initlib.DB.First(&note, "id = ?", nodeId)
	err = res.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with given ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Category != "" {
		updates["category"] = payload.Category
	}
	if payload.Content != "" {
		updates["content"] = payload.Content
	}
	if payload.Published != nil {
		updates["published"] = payload.Published
	}

	updates["updated_at"] = time.Now()
	initlib.DB.Model(&note).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": note}})
}

func SearchNoteByIdHandler(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	var note models.Note

	res := initlib.DB.First(&note, "id = ?", noteId)
	err := res.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with given ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": note}})
}

func DeleteNoteHandler(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	res := initlib.DB.Delete(&models.Note{}, "id = ?", noteId)

	if res.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with given ID exists"})
	} else if res.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": res.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}