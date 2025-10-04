package repository

import "laundry-backend/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	UserLogin(phoneNumber, password string) (*entity.User, error)
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

type UserCouponRepository interface {
	CreateUserCoupon(userCoupon *entity.UserCoupon) error
	GetByUserID(userID int) ([]entity.UserCoupon, error)
	GetByCouponID(couponID int) ([]entity.UserCoupon, error)
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
	GetByOrderID(orderID int) ([]entity.ClothList, error)
	Update(clothList *entity.ClothList) error
}

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	GetByUserID(userID int) ([]entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	Update(order *entity.Order) error
}