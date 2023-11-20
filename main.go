package main

import (
	"log"

	"embed"

	"github.com/arencloud/space-demo/controllers"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/soypat/rebed"
)

//go:embed docs/swagger.*
var swaggerFiles embed.FS



func main() {
	rebed.Write(swaggerFiles, "")
	cfg := swagger.Config{
		BasePath: "/",
        Path: "docs",
		FilePath: "docs/swagger.json",
	}
	user := controllers.New()
	app := fiber.New()
	mApp := fiber.New()
	app.Mount("/api", mApp)
	app.Use(logger.New())
	app.Use(swagger.New(cfg))

	mApp.Route("/users", func(router fiber.Router) {
		router.Post("/create", user.CreateUserHandler)
		router.Get("/list", user.FindUsersHandler)
	})
	mApp.Route("/users", func(router fiber.Router) {
		router.Delete("/delete/:id", user.DeleteUserHandler)
		router.Get("/list/:id", user.FindUserByIdHandler)
		router.Patch("/update/:id", user.UpdateUserHandler)
	})
	mApp.Get("/healz", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "healthy",
		})

	})
	log.Fatal(app.Listen(":8080"))
}
