package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
)

type CouponUsercase struct {
	couponRepo repository.CouponRepository
}

func NewCouponUsecase(couponRepo repository.CouponRepository) *CouponUsercase {
	return &CouponUsercase{couponRepo: couponRepo}
}

func (uc *CouponUsercase) CreateCoupon(coupon *entity.Coupon) error {
	return uc.couponRepo.CreateCoupon(coupon)
}

func (uc *CouponUsercase) GetCouponByID(id int) (*entity.Coupon, error) {
	return uc.couponRepo.GetByID(id)
}

func (uc *CouponUsercase) GetAllCoupons() ([]entity.Coupon, error) {
	return uc.couponRepo.GetAllCoupons()
}

func (uc *CouponUsercase) UpdateCoupon(coupon *entity.Coupon) error {
	return uc.couponRepo.Update(coupon)
}