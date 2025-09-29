package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserCouponHandler struct {
	userCouponUsecase usecase.UserCouponUsecase
}

func NewUserCouponHandler(userCouponUsecase usecase.UserCouponUsecase) *UserCouponHandler {
	return &UserCouponHandler{userCouponUsecase: userCouponUsecase}
}

func (h *UserCouponHandler) CreateUserCoupon(c *fiber.Ctx) error {
	var userCoupon entity.UserCoupon
	if err := c.BodyParser(&userCoupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := h.userCouponUsecase.CreateUserCoupon(&userCoupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(userCoupon)
}

func (h *UserCouponHandler) GetUserCouponsByUserID(c *fiber.Ctx) error {
	idParam := c.Params("user_id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}
	userCoupons, err := h.userCouponUsecase.GetUserCouponsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupons)
}

func (h *UserCouponHandler) GetUserCouponsByCouponID(c *fiber.Ctx) error {
	idParam := c.Params("coupon_id")
	couponID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid coupon ID"})
	}
	userCoupons, err := h.userCouponUsecase.GetUserCouponsByCouponID(couponID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupons)
}

func (h *UserCouponHandler) GetAllUserCoupons(c *fiber.Ctx) error {
	userCoupons, err := h.userCouponUsecase.GetAllUserCoupons()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupons)
}

func (h *UserCouponHandler) UpdateUserCoupon(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user coupon ID"})
	}
	var userCoupon entity.UserCoupon
	if err := c.BodyParser(&userCoupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	userCoupon.ID = uint(id)
	if err := h.userCouponUsecase.UpdateUserCoupon(&userCoupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupon)
}