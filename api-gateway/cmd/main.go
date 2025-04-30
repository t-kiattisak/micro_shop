package main

import (
	"api-gateway/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found. Using system environment variables.")
	}

	app := fiber.New()

	routes.SetupProxyRoutes(app)
	log.Println("Starting Order Service on port 8080...")
	app.Listen(":8080")
}
