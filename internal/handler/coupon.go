package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CouponHandler struct {
	couponUsecase usecase.CouponUsecase
}

func NewCouponHandler(couponUsecase usecase.CouponUsecase) *CouponHandler {
	return &CouponHandler{couponUsecase: couponUsecase}
}

func (h *CouponHandler) CreateCoupon(c *fiber.Ctx) error {
	var coupon entity.Coupon
	if err := c.BodyParser(&coupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if err := h.couponUsecase.CreateCoupon(&coupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(coupon)
}

func (h *CouponHandler) GetCouponByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID คูปองไม่ถูกต้อง"})
	}
	coupon, err := h.couponUsecase.GetCouponByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if coupon == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ไม่เจอคูปอง"})
	}
	return c.JSON(coupon)
}

func (h *CouponHandler) GetAllCoupons(c *fiber.Ctx) error {
	coupons, err := h.couponUsecase.GetAllCoupons()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(coupons)
}

func (h *CouponHandler) UpdateCoupon(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID คูปองไม่ถูกต้อง"})
	}
	var coupon entity.Coupon
	if err := c.BodyParser(&coupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	coupon.ID = id
	if err := h.couponUsecase.UpdateCoupon(&coupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(coupon)
}