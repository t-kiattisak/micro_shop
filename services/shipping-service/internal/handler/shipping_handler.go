package handler

import (
	"shipping-service/internal/usecase"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ShippingHandler struct {
	usecase *usecase.ShippingUseCase
}

func NewShippingHandler(uc *usecase.ShippingUseCase) *ShippingHandler {
	return &ShippingHandler{usecase: uc}
}

func (h *ShippingHandler) CreateShipping(c *fiber.Ctx) error {
	var req struct {
		OrderID        uint   `json:"order_id"`
		Carrier        string `json:"carrier"`
		TrackingNumber string `json:"tracking_number"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.usecase.CreateShipping(req.OrderID, req.Carrier, req.TrackingNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create shipping"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Shipping created successfully"})
}

func (h *ShippingHandler) UpdateShippingStatus(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	status := c.FormValue("status")
	orderIDUint, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}
	err = h.usecase.UpdateShippingStatus(uint(orderIDUint), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update status"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Shipping status updated"})
}
