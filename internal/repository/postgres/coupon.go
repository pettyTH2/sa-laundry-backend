package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"gorm.io/gorm"
)

type CouponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) repository.CouponRepository {
	return &CouponRepository{db: db}
}

func (r *CouponRepository) CreateCoupon(coupon *entity.Coupon) error {
	return r.db.Create(coupon).Error
}

func (r *CouponRepository) GetByID(id int) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.First(&coupon, id).Error
	return &coupon, err
}

func (r *CouponRepository) GetAllCoupons() ([]entity.Coupon, error) {
	var coupons []entity.Coupon
	err := r.db.Find(&coupons).Error
	return coupons, err
}

func (r *CouponRepository) Update(coupon *entity.Coupon) error {
	return r.db.Save(coupon).Error
}