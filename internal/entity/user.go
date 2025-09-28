package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" gorm:"unique"`
	Name	 string `json:"name"`
	Nickname	 string `json:"nickname"`
	Password string `json:"password"`
	Role	 string `json:"role"`

	UserCoupons []UserCoupon `json:"user_coupons" gorm:"foreignKey:UserID"`
	Orders       []Order       `json:"orders" gorm:"foreignKey:UserID"`
}