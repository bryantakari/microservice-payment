package app

import (
	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/bryantakari/microservice-payment/order-service/internal/order"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(conf *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "order-service",
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Health check
	orderHandler := order.Handler{}
	orderHandler.RegisterRoutes(app)
	return app
}
