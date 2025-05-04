package handler

import (
	"github.com/gofiber/fiber/v2"
)

type MeHandler struct{}

func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

func (h *MeHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	role := c.Locals("role")

	return c.JSON(fiber.Map{
		"user_id": userID,
		"role":    role,
	})
}
