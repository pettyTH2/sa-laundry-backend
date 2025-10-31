package entity

import (
	"github.com/google/uuid"
)

type ClothList struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Quantity     int            `json:"quantity" gorm:"not null"`
	SubTotalCost int            `json:"sub_total_cost" gorm:"not null"`
	OrderID      uuid.UUID      `json:"order_id" gorm:"type:uuid;not null"`
	ClothID      uuid.UUID      `json:"cloth_id" gorm:"type:uuid;not null"`
	
	Order Order `json:"order" gorm:"foreignKey:OrderID"`
	Cloth Cloth `json:"cloth" gorm:"foreignKey:ClothID"`
}

func (ClothList) TableName() string {
	return "cloth_list"
}