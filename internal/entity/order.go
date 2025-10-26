package entity

import (
	"gorm.io/gorm"
)	

type Order struct {
	gorm.Model
	ServiceType string `json:"service_type" gorm:"not null"`
	TotalCloth int `json:"total_cloth" gorm:"not null"`
	TotalCost int `json:"total_cost" gorm:"not null"`
	OrderDate string `json:"order_date"`
	PickupDate string `json:"pickup_date"`
	OrderStatus string `json:"order_status" gorm:"not null"`
	PaymentMethod string `json:"payment_method" gorm:"not null"`
	UserID    int  `json:"user_id" gorm:"not null"`

	User User   `json:"user" gorm:"foreignKey:UserID"`
	ClothList []ClothList `json:"cloth_list" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "order"
}