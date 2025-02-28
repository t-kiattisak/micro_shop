package handler

import (
	"order-service/internal/domain"
	"order-service/internal/dto"
	"order-service/internal/usecase"
	"strconv"

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
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	if err := h.usecase.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Err(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Ok("Order created successfully", order))
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orders, err := h.usecase.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Err(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(dto.Ok("Order get successfully", orders))
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err("Invalid order ID"))
	}

	order, err := h.usecase.GetOrderByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.Err(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.Ok("Order get successfully", order))
}

func (h *OrderHandler) DeleteOrderByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err("Invalid order ID"))
	}

	order, err := h.usecase.DeleteOrderByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.Err(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.Ok("Order Delete successfully", order))
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	newStatus := c.FormValue("status")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err("Invalid order ID"))
	}

	if err := h.usecase.UpdateOrderStatus(uint(id), newStatus); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.Ok("Order status updated successfully", nil))
}
