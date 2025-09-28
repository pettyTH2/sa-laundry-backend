package entity

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	CPName string `json:"cp_name"`
	CPPrice int `json:"cp_price"`

	UserCoupons []UserCoupon `json:"user_coupons" gorm:"foreignKey:CouponID"`
}