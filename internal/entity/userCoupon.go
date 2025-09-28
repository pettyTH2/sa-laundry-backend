package entity

import "gorm.io/gorm"

type UserCoupon struct {
	gorm.Model
	PointLeft int   `json:"point_left"`
	StartDate string `json:"start_date"`
	ExpireDate string `json:"expire_date"`
	UserID    int  `json:"user_id"`
	CouponID int  `json:"coupon_id"`

	User   User   `json:"users" gorm:"foreignKey:UserID"`
    Coupon Coupon `json:"coupons" gorm:"foreignKey:CouponID"`
}