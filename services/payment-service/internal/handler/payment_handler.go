package handler

import (
	"payment-service/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	usecase *usecase.PaymentUseCase
}

func NewPaymentHandler(uc *usecase.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{usecase: uc}
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	type Request struct {
		OrderID uint    `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	payment, err := h.usecase.CreatePayment(req.OrderID, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(payment)
}

func (h *PaymentHandler) GetPayment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	payment, err := h.usecase.GetPayment(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "payment not found"})
	}

	return c.JSON(payment)
}
