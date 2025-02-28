package main

import (
	"log"

	"order-service/infrastructure"
	"order-service/internal"
	"order-service/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := infrastructure.ConnectDB()
	db.Exec("SELECT 1")

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Order{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	orderHandler := internal.CreateOrderHandler()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Order Service Running"})
	})

	app.Post("/orders", orderHandler.CreateOrder)
	app.Get("/orders", orderHandler.GetOrders)
	app.Get("/orders/:id", orderHandler.GetOrderByID)
	app.Delete("/orders/:id", orderHandler.DeleteOrderByID)
	app.Patch("/orders/:id/status", orderHandler.UpdateOrderStatus)

	log.Println("Starting Order Service on port 8081...")
	log.Fatal(app.Listen(":8081"))
}
