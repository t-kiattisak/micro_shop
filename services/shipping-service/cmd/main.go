package main

import (
	"log"
	"shipping-service/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	shippingHandler := internal.CreateShippingHandler()

	app.Post("/shopping", shippingHandler.CreateShipping)
	app.Patch("/shopping/:order_id/status", shippingHandler.UpdateShippingStatus)

	log.Println("✅ Shipping Service is running on port 8086...")
	log.Fatal(app.Listen(":8086"))
}
