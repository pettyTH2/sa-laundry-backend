package entity

import (
	"github.com/google/uuid"
)

type Coupon struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CPName    string         `json:"cp_name" gorm:"not null"`
	CPPrice   int            `json:"cp_price" gorm:"not null"`
	
	UserCoupons []UserCoupon `json:"user_coupons" gorm:"foreignKey:CouponID"`
}

func (Coupon) TableName() string {
	return "coupon"
}