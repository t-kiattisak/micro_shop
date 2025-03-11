package handler

import (
	"encoding/json"
	"log"
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

func (h *PaymentHandler) PaymentWebhook(c *fiber.Ctx) error {
	var payload struct {
		OrderID uint   `json:"order_id"`
		Status  string `json:"status"`
	}

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	log.Printf("🔔 Received Webhook: Order %d - Status: %s\n", payload.OrderID, payload.Status)

	if payload.Status == "PAID" {
		err := h.usecase.UpdatePaymentStatus(payload.OrderID, "PAID")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {
		err := h.usecase.UpdatePaymentStatus(payload.OrderID, "FAILED")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return c.JSON(fiber.Map{"message": "Webhook processed"})
}
