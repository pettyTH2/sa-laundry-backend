package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"github.com/google/uuid"
)

type CouponUsecase struct {
	couponRepo repository.CouponRepository
}

func NewCouponUsecase(couponRepo repository.CouponRepository) *CouponUsecase {
	return &CouponUsecase{couponRepo: couponRepo}
}

func (uc *CouponUsecase) CreateCoupon(coupon *entity.Coupon) error {
	return uc.couponRepo.CreateCoupon(coupon)
}

func (uc *CouponUsecase) GetCouponByID(id uuid.UUID) (*entity.Coupon, error) {
	return uc.couponRepo.GetByID(id)
}

func (uc *CouponUsecase) GetAllCoupons() ([]entity.Coupon, error) {
	return uc.couponRepo.GetAllCoupons()
}

func (uc *CouponUsecase) UpdateCoupon(coupon *entity.Coupon) error {
	return uc.couponRepo.Update(coupon)
}