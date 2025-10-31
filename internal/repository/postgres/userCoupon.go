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
	err := r.db.Table("user_coupon").
		Select("id, point_left, start_date, expire_date, user_id, coupon_id").
		Where("user_id = ?", userID).
		Scan(&userCoupons).Error

	if err != nil {
		return userCoupons, err
	}

	// Load User and Coupon for each UserCoupon
	for i := range userCoupons {
		r.db.Table("user").
			Select("id, phone_number, name, nickname, role").
			Where("id = ?", userCoupons[i].UserID).
			Scan(&userCoupons[i].User)

		r.db.Table("coupon").
			Select("id, cp_name, cp_price").
			Where("id = ?", userCoupons[i].CouponID).
			Scan(&userCoupons[i].Coupon)
	}

	return userCoupons, nil
}

func (r *UserCouponRepository) GetByCouponID(couponID uuid.UUID) ([]entity.UserCoupon, error) {
	var userCoupons []entity.UserCoupon
	err := r.db.Table("user_coupon").
		Select("id, point_left, start_date, expire_date, user_id, coupon_id").
		Where("coupon_id = ?", couponID).
		Scan(&userCoupons).Error

	if err != nil {
		return userCoupons, err
	}

	// Load User and Coupon for each UserCoupon
	for i := range userCoupons {
		r.db.Table("user").
			Select("id, phone_number, name, nickname, role").
			Where("id = ?", userCoupons[i].UserID).
			Scan(&userCoupons[i].User)

		r.db.Table("coupon").
			Select("id, cp_name, cp_price").
			Where("id = ?", userCoupons[i].CouponID).
			Scan(&userCoupons[i].Coupon)
	}

	return userCoupons, nil
}

func (r *UserCouponRepository) GetAllUserCoupons() ([]entity.UserCoupon, error) {
	var userCoupons []entity.UserCoupon
	err := r.db.Table("user_coupon").
		Select("id, point_left, start_date, expire_date, user_id, coupon_id").
		Scan(&userCoupons).Error

	if err != nil {
		return userCoupons, err
	}

	// Load User and Coupon for each UserCoupon
	for i := range userCoupons {
		r.db.Table("user").
			Select("id, phone_number, name, nickname, role").
			Where("id = ?", userCoupons[i].UserID).
			Scan(&userCoupons[i].User)

		r.db.Table("coupon").
			Select("id, cp_name, cp_price").
			Where("id = ?", userCoupons[i].CouponID).
			Scan(&userCoupons[i].Coupon)
	}

	return userCoupons, nil
}

func (r *UserCouponRepository) Update(userCoupon *entity.UserCoupon) error {
	return r.db.Table("user_coupon").
		Where("id = ?", userCoupon.ID).
		Updates(map[string]interface{}{
			"point_left":  userCoupon.PointLeft,
			"start_date":  userCoupon.StartDate,
			"expire_date": userCoupon.ExpireDate,
			"user_id":     userCoupon.UserID,
			"coupon_id":   userCoupon.CouponID,
		}).Error
}
