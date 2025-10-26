package entity

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	CPName string `json:"cp_name" gorm:"not null"`
	CPPrice int `json:"cp_price" gorm:"not null"`

	UserCoupons []UserCoupon `json:"user_coupons" gorm:"foreignKey:CouponID"`
}

func (Coupon) TableName() string {
	return "coupon"
}