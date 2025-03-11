package main

import (
	"log"
	"payment-service/infrastructure"
	"payment-service/internal"
	"payment-service/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := infrastructure.ConnectDB()
	db.Exec("SELECT 1")

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Payment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	paymentHandler := internal.CreatePaymentHandler()

	app.Get("/payments/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Payment Service Running"})
	})

	app.Post("/payments", paymentHandler.CreatePayment)
	app.Get("/payments/:id", paymentHandler.GetPayment)
	app.Post("/payments/webhook", paymentHandler.PaymentWebhook)

	log.Println("Starting Order Service on port 8081...")
	log.Fatal(app.Listen(":8083"))
}
