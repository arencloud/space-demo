package main

import (
	"log"

	"github.com/arencloud/space-demo/controllers"
	"github.com/arencloud/space-demo/initlib"
	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config := initlib.LoadConfig()
	initlib.ConnectDatabase(&config)
}

func main() {
	app := fiber.New()
	mApp := fiber.New()
	app.Mount("/api", mApp)
	app.Use(logger.New())

	mApp.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.SearchNotesHandler)
	})
	mApp.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteNoteHandler)
		router.Get("", controllers.SearchNoteByIdHandler)
		router.Patch("", controllers.UpdateNoteHandler)
	})
	mApp.Get("/healz", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"message": "healthy",
		})
		
	})
	log.Fatal(app.Listen(":8080"))
}