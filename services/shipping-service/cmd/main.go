package main

import (
	"log"
	"shipping-service/infrastructure"
	"shipping-service/internal/domain"
	"shipping-service/internal/handler"
	"shipping-service/internal/repository"
	"shipping-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := infrastructure.ConnectDB()

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Shipping{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	productRepo := repository.NewShippingRepository(db)
	productUseCase := usecase.NewShippingUseCase(productRepo)
	productHandler := handler.NewShippingHandler(productUseCase)

	app.Post("/shopping", productHandler.CreateShipping)
	app.Patch("/shopping/:order_id/status", productHandler.UpdateShippingStatus)

	log.Println("✅ Shipping Service is running on port 8086...")
	log.Fatal(app.Listen(":8086"))
}
