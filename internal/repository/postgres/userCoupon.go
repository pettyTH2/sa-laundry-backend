package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)	

type UserCouponRepository struct {
	db *gorm.DB
}

func NewUserCouponRepository(db *gorm.DB) repository.UserCouponRepository {
	return &UserCouponRepository{db: db}
}

func (r *UserCouponRepository) CreateUserCoupon(userCoupon *entity.UserCoupon) error {
	return r.db.Create(userCoupon).Error
}

func (r *UserCouponRepository) GetByUserID(userID uuid.UUID) ([]entity.UserCoupon, error) {
	var userCoupons []entity.UserCoupon
	err := r.db.Preload("User").Preload("Coupon").Where("user_id = ?", userID).Find(&userCoupons).Error
	return userCoupons, err
}

func (r *UserCouponRepository) GetByCouponID(couponID uuid.UUID) ([]entity.UserCoupon, error) {
	var userCoupons []entity.UserCoupon
	err := r.db.Preload("User").Preload("Coupon").Where("coupon_id = ?", couponID).Find(&userCoupons).Error
	return userCoupons, err
}

func (r *UserCouponRepository) GetAllUserCoupons() ([]entity.UserCoupon, error) {
	var userCoupons []entity.UserCoupon
	err := r.db.Preload("User").Preload("Coupon").Find(&userCoupons).Error
	return userCoupons, err
}

func (r *UserCouponRepository) Update(userCoupon *entity.UserCoupon) error {
	return r.db.Save(userCoupon).Error
}