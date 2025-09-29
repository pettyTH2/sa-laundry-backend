package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type CouponHandler struct {
	couponUsecase usecase.CouponUsercase
}

func NewCouponHandler(couponUsecase usecase.CouponUsercase) *CouponHandler {
	return &CouponHandler{couponUsecase: couponUsecase}
}

func (h *CouponHandler) CreateCoupon(c *fiber.Ctx) error {
	var coupon entity.Coupon
	if err := c.BodyParser(&coupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := h.couponUsecase.CreateCoupon(&coupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(coupon)
}

func (h *CouponHandler) GetCouponByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid coupon ID"})
	}
	coupon, err := h.couponUsecase.GetCouponByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if coupon == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "coupon not found"})
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
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid coupon ID"})
	}
	var coupon entity.Coupon
	if err := c.BodyParser(&coupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	coupon.ID = uint(id)
	if err := h.couponUsecase.UpdateCoupon(&coupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(coupon)
}