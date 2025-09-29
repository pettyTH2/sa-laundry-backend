package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
)

type UserCouponUsecase struct {
	userCouponRepo repository.UserCouponRepository
}

func NewUserCouponUsecase(userCouponRepo repository.UserCouponRepository) *UserCouponUsecase {
	return &UserCouponUsecase{userCouponRepo: userCouponRepo}
}

func (uc *UserCouponUsecase) CreateUserCoupon(userCoupon *entity.UserCoupon) error {
	return uc.userCouponRepo.CreateUserCoupon(userCoupon)
}

func (uc *UserCouponUsecase) GetUserCouponsByUserID(userID int) ([]entity.UserCoupon, error) {
	return uc.userCouponRepo.GetByUserID(userID)
}

func (uc *UserCouponUsecase) GetUserCouponsByCouponID(couponID int) ([]entity.UserCoupon, error) {
	return uc.userCouponRepo.GetByCouponID(couponID)
}

func (uc *UserCouponUsecase) GetAllUserCoupons() ([]entity.UserCoupon, error) {
	return uc.userCouponRepo.GetAllUserCoupons()
}

func (uc *UserCouponUsecase) UpdateUserCoupon(userCoupon *entity.UserCoupon) error {
	return uc.userCouponRepo.Update(userCoupon)
}