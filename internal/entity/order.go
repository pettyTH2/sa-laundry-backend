package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ServiceType string `json:"service_type"`
	TotalCloth int `json:"total_cloth"`
	TotalCost int `json:"total_cost"`
	OrderDate string `json:"order_date"`
	PickupDate string `json:"pickup_date"`
	OrderStatus string `json:"orders_status"`
	PaymentMethod string `json:"payment_method"`
	UserID    int  `json:"user_id"`

	User User   `json:"user" gorm:"foreignKey:UserID"`
	ClothList []ClothList `json:"cloth_list" gorm:"foreignKey:OrderID"`
}