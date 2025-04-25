package main

import (
	"api-gateway/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupProxyRoutes(app)
	log.Println("Starting Order Service on port 8080...")
	app.Listen(":8080")
}
