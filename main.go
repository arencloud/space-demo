package main

import (
	"log"

	"github.com/arencloud/space-demo/controllers"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)



func main() {
	cfg := swagger.Config{
		BasePath: "/",
        Path: "docs",
		FilePath: "docs/swagger.json",
	}
	user := controllers.New()
	app := fiber.New()
	api := fiber.New()
	mApp := fiber.New()
	app.Mount("/", api)
	app.Use(logger.New())
	app.Use(swagger.New(cfg))

	//app.Static("", "public")

	api.Mount("api", mApp)

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
