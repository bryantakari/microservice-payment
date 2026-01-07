package app

import (
	"log/slog"

	"github.com/bryantakari/microservice-payment/order-service/internal/order"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(orderService order.Service, log *slog.Logger) *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:               "order-service",
		DisableStartupMessage: true,
	})
	// Global middleware
	app.Use(recover.New())

	orderHandler := order.NewHandler(orderService, log)
	orderHandler.RegisterRoutes(app)
	log.Info("order-service ready")
	return app
}
