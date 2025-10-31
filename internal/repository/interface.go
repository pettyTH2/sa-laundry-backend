package repository

import (
	"laundry-backend/internal/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	UserLogin(phoneNumber, password string) (*entity.User, error)
	GetByRole(role string) ([]entity.User, error)
	GetById(id string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	Update(user *entity.User) error
}

type CouponRepository interface {
	CreateCoupon(coupon *entity.Coupon) error
	GetByID(id uuid.UUID) (*entity.Coupon, error)
	GetAllCoupons() ([]entity.Coupon, error)
	Update(coupon *entity.Coupon) error
}

type UserCouponRepository interface {
	CreateUserCoupon(userCoupon *entity.UserCoupon) error
	GetByUserID(userID uuid.UUID) ([]entity.UserCoupon, error)
	GetByCouponID(couponID uuid.UUID) ([]entity.UserCoupon, error)
	GetAllUserCoupons() ([]entity.UserCoupon, error)
	Update(userCoupon *entity.UserCoupon) error
}

type ClothRepository interface {
	CreateCloth(cloth *entity.Cloth) error
	GetByCategory(catagory string) ([]entity.Cloth, error)
	GetAllCloths() ([]entity.Cloth, error)
	Update(cloth *entity.Cloth) error
}

type ClothListRepository interface {
	CreateClothList(clothList *entity.ClothList) error
	GetByOrderID(orderID uuid.UUID) ([]entity.ClothList, error)
	Update(clothList *entity.ClothList) error
}

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	GetByUserID(userID uuid.UUID) ([]entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	Update(order *entity.Order) error
	GetByID(id uuid.UUID) (*entity.Order, error)
}