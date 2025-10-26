package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" gorm:"unique;not null"`
	Name	 string `json:"name" gorm:"not null"`
	Nickname	 string `json:"nickname" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Role	 string `json:"role" gorm:"not null"`

	UserCoupons []UserCoupon `json:"user_coupons" gorm:"foreignKey:UserID"`
	Orders       []Order       `json:"orders" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}