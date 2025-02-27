package handler

import (
	"order-service/internal/domain"
	"order-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	usecase *usecase.OrderUseCase
}

func NewOrderHandler(uc *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{usecase: uc}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order domain.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.usecase.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
