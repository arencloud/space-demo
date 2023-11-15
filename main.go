package main

import (
	"log"

	"github.com/arencloud/space-demo/controllers"
	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	user := controllers.New()
	app := fiber.New()
	mApp := fiber.New()
	app.Mount("/api", mApp)
	app.Use(logger.New())

	mApp.Route("/notes", func(router fiber.Router) {
		router.Post("/", user.CreateUserHandler)
		router.Get("", user.FindUsersHandler)
	})
	mApp.Route("/notes/:id", func(router fiber.Router) {
		router.Delete("", user.DeleteUserHandler)
		router.Get("", user.FindUserByIdHandler)
		router.Patch("", user.UpdateUserHandler)
	})
	mApp.Get("/healz", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "healthy",
		})

	})
	log.Fatal(app.Listen(":8080"))
}
