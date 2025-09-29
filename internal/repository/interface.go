package repository

import "laundry-backend/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetByRole(role string) ([]entity.User, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	Update(user *entity.User) error
}

type CouponRepository interface {
	CreateCoupon(coupon *entity.Coupon) error
	GetByID(id int) (*entity.Coupon, error)
	GetAllCoupons() ([]entity.Coupon, error)
	Update(coupon *entity.Coupon) error
}