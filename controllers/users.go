package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/arencloud/space-demo/initlib"
	"github.com/arencloud/space-demo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandlers struct {
	db *gorm.DB
}

func New() *UserHandlers {
	db := initlib.InitDb()
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{})
	return &UserHandlers{
		db: db,
	}
}

func (h *UserHandlers) CreateUserHandler(c *fiber.Ctx) error {
	var payload models.User

	json.Unmarshal(c.Body(), &payload)

	fmt.Println(payload)

	result := models.CreateUser(h.db, &payload)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": payload}})
}

func (h *UserHandlers) FindUsersHandler(c *fiber.Ctx) error {
	var users []models.User

	result := models.GetUsers(h.db, &users)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(users), "users": users})
}

func (h *UserHandlers) UpdateUserHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User
	id, _ := strconv.Atoi(userId)

	result := models.GetUser(h.db, &user, id)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with given ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	json.Unmarshal(c.Body(), &user)

	_ = models.UpdateUser(h.db, &user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func (h *UserHandlers) FindUserByIdHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User

	id, _ := strconv.Atoi(userId)

	result := models.GetUser(h.db, &user, id)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with given ID exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func (h *UserHandlers) DeleteUserHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User
	id, _ := strconv.Atoi(userId)
	result := models.DeleteUser(h.db, &user, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No user with given ID exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
