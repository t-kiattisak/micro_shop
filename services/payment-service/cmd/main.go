package main

import (
	"log"
	"payment-service/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

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
