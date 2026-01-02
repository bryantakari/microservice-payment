package order

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	srv Service
	log *slog.Logger
}

var _ = (*Handler)(nil)

func NewHandler(srv Service, log *slog.Logger) *Handler {
	return &Handler{
		srv: srv,
		log: log.With("layer", "handler"),
	}

}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Get("/health", h.healthCheck)
}

func (h *Handler) healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
