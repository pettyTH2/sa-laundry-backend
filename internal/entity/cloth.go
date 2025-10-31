package entity

import (
	"github.com/google/uuid"
)

type Cloth struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ClothName  string         `json:"cloth_name" gorm:"not null"`
	ClothPrice int            `json:"cloth_price" gorm:"not null"`
	Category   string         `json:"category" gorm:"not null"`
	
	ClothLists []ClothList `json:"cloth_lists" gorm:"foreignKey:ClothID"`
}

func (Cloth) TableName() string {
	return "cloth"
}