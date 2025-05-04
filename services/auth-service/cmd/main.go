package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadEnv()

	app := fiber.New()
	h := handler.NewAuthHandler(cfg.JWTSecret)
	meHandler := handler.NewMeHandler()
	authRoute := app.Group("/auth")

	authRoute.Post("/login", h.Login)
	authRoute.Get("/me", middlewares.JWTMiddleware(cfg.JWTSecret), meHandler.Me)

	log.Printf("Starting Auth Service on port %s...", cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}
