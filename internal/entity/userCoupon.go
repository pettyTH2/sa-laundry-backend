package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserCoupon struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PointLeft  int       `json:"point_left"`
	StartDate  time.Time `json:"start_date" gorm:"type:timestamp"`
	ExpireDate time.Time `json:"expire_date" gorm:"type:timestamp"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CouponID   uuid.UUID `json:"coupon_id" gorm:"type:uuid;not null"`

	User   User   `json:"users" gorm:"foreignKey:UserID"`
	Coupon Coupon `json:"coupons" gorm:"foreignKey:CouponID"`
}

func (UserCoupon) TableName() string {
	return "user_coupon"
}
