package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "order-service",
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}
