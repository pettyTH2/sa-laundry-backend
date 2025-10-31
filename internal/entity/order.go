package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ServiceType   string    `json:"service_type" gorm:"not null"`
	TotalCloth    int       `json:"total_cloth" gorm:"not null"`
	TotalCost     int       `json:"total_cost" gorm:"not null"`
	OrderDate     time.Time `json:"order_date" gorm:"type:timestamp"`
	PickupDate    time.Time `json:"pickup_date" gorm:"type:timestamp"`
	OrderStatus   string    `json:"order_status" gorm:"not null"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`

	User       User        `json:"user" gorm:"foreignKey:UserID"`
	ClothLists []ClothList `json:"cloth_lists" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "order"
}
