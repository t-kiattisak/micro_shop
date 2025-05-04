package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadEnv()

	app := fiber.New()
	h := handler.NewAuthHandler(cfg.JWTSecret)

	app.Post("/login", h.Login)
	app.Listen(cfg.Port)
}
