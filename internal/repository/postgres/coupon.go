package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"

	"github.com/google/uuid"
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

func (r *CouponRepository) GetByID(id uuid.UUID) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.Table("coupon").
		Select("id, cp_name, cp_price").
		Where("id = ?", id).
		Limit(1).
		Scan(&coupon).Error
	return &coupon, err
}

func (r *CouponRepository) GetAllCoupons() ([]entity.Coupon, error) {
	var coupons []entity.Coupon
	err := r.db.Table("coupon").
		Select("id, cp_name, cp_price").
		Scan(&coupons).Error
	return coupons, err
}

func (r *CouponRepository) Update(coupon *entity.Coupon) error {
	return r.db.Table("coupon").
		Where("id = ?", coupon.ID).
		Updates(map[string]interface{}{
			"cp_name":  coupon.CPName,
			"cp_price": coupon.CPPrice,
		}).Error
}
