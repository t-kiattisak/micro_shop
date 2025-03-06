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

func (h *InventoryHandler) CreateInventory(c *fiber.Ctx) error {
	type Request struct {
		Product  string `json:"product"`
		Quantity int    `json:"quantity"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	err := h.usecase.CreateInventory(req.Product, req.Quantity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Err(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Ok("inventory created", nil))
}
