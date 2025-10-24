package handler	

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type ClothListHandler struct {
	clothListUsecase usecase.ClothListUsecase
}

func NewClothListHandler(clothListUsecase usecase.ClothListUsecase) *ClothListHandler {
	return &ClothListHandler{clothListUsecase: clothListUsecase}
}

func (h *ClothListHandler) CreateClothList(c *fiber.Ctx) error {
	var clothList entity.ClothList
	if err := c.BodyParser(&clothList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if err := h.clothListUsecase.CreateClothList(&clothList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(clothList)
}

func (h *ClothListHandler) GetClothListsByOrderID(c *fiber.Ctx) error {
	idParam := c.Params("order_id")
	orderID, err := strconv.Atoi(idParam)	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID ออร์เดอร์ไม่ถูกต้อง"})
	}
	clothLists, err := h.clothListUsecase.GetClothListsByOrderID(orderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clothLists)
}

func (h *ClothListHandler) UpdateClothList(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID รายการเสื้อไม่ถูกต้อง"})
	}
	var clothList entity.ClothList
	if err := c.BodyParser(&clothList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	clothList.ID = uint(id)
	if err := h.clothListUsecase.UpdateClothList(&clothList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(clothList)
}