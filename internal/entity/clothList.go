package entity

import "gorm.io/gorm"

type ClothList struct {
	gorm.Model
	Quantity int `json:"quantity" gorm:"not null"`
	SubTotalCost int `json:"sub_total_cost" gorm:"not null"`
	OrderID int `json:"order_id" gorm:"not null"`
	ClothID int `json:"cloth_id" gorm:"not null"`

	Order Order `json:"order" gorm:"foreignKey:OrderID"`
	Cloth Cloth `json:"cloth" gorm:"foreignKey:ClothID"`
}

func (ClothList) TableName() string {
	return "cloth_list"
}