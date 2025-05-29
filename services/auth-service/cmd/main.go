package main

import (
	"auth-service/infrastructure"
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middlewares"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadEnv()
	db := infrastructure.ConnectDB()

	app := fiber.New()

	repo := repository.NewUserRepository(db)
	u := usecase.NewAuthUsecase(repo)
	h := handler.NewAuthHandler(u)

	meHandler := handler.NewMeHandler()
	authRoute := app.Group("/auth")

	authRoute.Post("/register", h.Register)

	authRoute.Post("/login", h.Login)
	authRoute.Get("/me", middlewares.JWTMiddleware(cfg.JWTSecret), meHandler.Me)

	log.Printf("Starting Auth Service on port %s...", cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}
