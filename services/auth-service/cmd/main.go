package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadEnv()

	app := fiber.New()
	h := handler.NewAuthHandler(cfg.JWTSecret)

	app.Post("/auth/login", h.Login)

	log.Printf("Starting Order Service on port %s...", cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}
