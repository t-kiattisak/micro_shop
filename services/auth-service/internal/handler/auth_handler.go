package handler

import (
	"auth-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := h.usecase.Register(req.Username, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}
