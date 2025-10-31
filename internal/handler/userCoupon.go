package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserCouponHandler struct {
	userCouponUsecase usecase.UserCouponUsecase
}

func NewUserCouponHandler(userCouponUsecase usecase.UserCouponUsecase) *UserCouponHandler {
	return &UserCouponHandler{userCouponUsecase: userCouponUsecase}
}

func (h *UserCouponHandler) CreateUserCoupon(c *fiber.Ctx) error {
	var req struct {
		PointLeft  int       `json:"point_left"`
		StartDate  time.Time `json:"start_date"`
		ExpireDate time.Time `json:"expire_date"`
		UserID     string    `json:"user_id"`
		CouponID   string    `json:"coupon_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	// Parse UUIDs
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID ไม่ถูกต้อง"})
	}

	couponID, err := uuid.Parse(req.CouponID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Coupon ID ไม่ถูกต้อง"})
	}

	// Create UserCoupon entity
	userCoupon := entity.UserCoupon{
		PointLeft:  req.PointLeft,
		StartDate:  req.StartDate,
		ExpireDate: req.ExpireDate,
		UserID:     userID,
		CouponID:   couponID,
	}

	// Set default values if not provided
	if userCoupon.PointLeft == 0 {
		userCoupon.PointLeft = 50
	}
	if userCoupon.StartDate.IsZero() {
		userCoupon.StartDate = time.Now()
	}
	if userCoupon.ExpireDate.IsZero() {
		userCoupon.ExpireDate = time.Now().AddDate(0, 1, 0)
	}

	if err := h.userCouponUsecase.CreateUserCoupon(&userCoupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(userCoupon)
}

func (h *UserCouponHandler) GetUserCouponsByUserID(c *fiber.Ctx) error {
	idParam := c.Params("user_id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID ผู้ใช้ไม่ถูกต้อง"})
	}
	userCoupons, err := h.userCouponUsecase.GetUserCouponsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupons)
}

func (h *UserCouponHandler) GetUserCouponsByCouponID(c *fiber.Ctx) error {
	idParam := c.Params("coupon_id")
	couponID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID คูปองไม่ถูกต้อง"})
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
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID คูปองของผู้ใช้ไม่ถูกต้อง"})
	}
	var userCoupon entity.UserCoupon
	if err := c.BodyParser(&userCoupon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	userCoupon.ID = id
	if err := h.userCouponUsecase.UpdateUserCoupon(&userCoupon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(userCoupon)
}
