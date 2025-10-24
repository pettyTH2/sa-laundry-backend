package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"time"
	"strconv"
)

type OrderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	order.OrderDate = time.Now().Format("02/01/2006")
	order.PickupDate = time.Now().AddDate(0, 0, 3).Format("02/01/2006")

	if err := h.orderUsecase.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) GetOrdersByUserID(c *fiber.Ctx) error {
	idParam := c.Params("user_id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID สมาชิกไม่ถูกต้อง"})
	}
	orders, err := h.orderUsecase.GetOrdersByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.orderUsecase.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID ออร์เดอร์ไม่ถูกต้อง"})
	}
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	order.ID = uint(id)
	if err := h.orderUsecase.UpdateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(order)
}