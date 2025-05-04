package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	secret []byte
}

func NewAuthHandler(secret string) *AuthHandler {
	return &AuthHandler{
		secret: []byte(secret),
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	if body.Email != "user@example.com" || body.Password != "123456" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "user-id-123",
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenStr, err := token.SignedString(h.secret)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not sign token")
	}

	return c.JSON(fiber.Map{
		"access_token": tokenStr,
		"token_type":   "Bearer",
		"expires_in":   3600,
	})
}
