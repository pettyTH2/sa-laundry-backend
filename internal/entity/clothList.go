package entity

import "gorm.io/gorm"

type ClothList struct {
	gorm.Model
	Quantity int `json:"quantity"`
	SubTotalCost int `json:"sub_total_cost"`
	OrderID int `json:"order_id"`
	ClothID int `json:"cloth_id"`

	Order Order `json:"order" gorm:"foreignKey:OrderID"`
	Cloth Cloth `json:"cloth" gorm:"foreignKey:ClothID"`
}