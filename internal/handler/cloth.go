package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type ClothHandler struct {
	clothUsecase usecase.ClothUsecase
}

func NewClothHandler(clothUsecase usecase.ClothUsecase) *ClothHandler {
	return &ClothHandler{clothUsecase: clothUsecase}
}

func (h *ClothHandler) CreateCloth(c *fiber.Ctx) error {
	var cloth entity.Cloth
	if err := c.BodyParser(&cloth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := h.clothUsecase.CreateCloth(&cloth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(cloth)
}

func (h *ClothHandler) GetClothsByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	cloths, err := h.clothUsecase.GetClothsByCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cloths)
}

func (h *ClothHandler) GetAllCloths(c *fiber.Ctx) error {
	cloths, err := h.clothUsecase.GetAllCloths()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cloths)
}

func (h *ClothHandler) UpdateCloth(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid cloth ID"})
	}
	var cloth entity.Cloth
	if err := c.BodyParser(&cloth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	cloth.ID = uint(id)
	if err := h.clothUsecase.UpdateCloth(&cloth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cloth)
}