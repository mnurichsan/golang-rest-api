package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mnurichsan/golang-rest-api/handlers"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	v1 := app.Group("/v1")
	api := v1.Group("/api")
	api.Get("/user", handlers.GetUserIndex)
	api.Post("/user/create", handlers.CreateUser)

	api.Post("auth/login", handlers.Login)
}
