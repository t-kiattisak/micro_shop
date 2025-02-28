package handler

import (
	"inventory-service/internal/dto"
	"inventory-service/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type InventoryHandler struct {
	usecase *usecase.InventoryUseCase
}

func NewInventoryHandler(uc *usecase.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{usecase: uc}
}

func (h *InventoryHandler) CheckStock(c *fiber.Ctx) error {
	product := c.Query("product")
	qty := c.QueryInt("qty")

	err := h.usecase.CheckStock(product, qty)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	return c.JSON(dto.Ok("Stock available", nil))
}

func (h *InventoryHandler) ReduceStock(c *fiber.Ctx) error {
	product := c.FormValue("product")
	qtyStr := c.FormValue("qty")

	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	err = h.usecase.ReduceStock(product, qty)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	return c.JSON(dto.Ok("Stock reduced successfully", nil))
}
