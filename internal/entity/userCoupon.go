package entity

import "gorm.io/gorm"

type UserCoupon struct {
	gorm.Model
	PointLeft int   `json:"point_left"`
	StartDate string `json:"start_date"`
	ExpireDate string `json:"expire_date"`
	UserID    int  `json:"user_id" gorm:"not null"`
	CouponID int  `json:"coupon_id" gorm:"not null"`

	User   User   `json:"users" gorm:"foreignKey:UserID"`
    Coupon Coupon `json:"coupons" gorm:"foreignKey:CouponID"`
}

func (UserCoupon) TableName() string {
	return "user_coupon"
}