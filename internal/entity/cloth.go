package entity

import "gorm.io/gorm"

type Cloth struct {
	gorm.Model
	ClothName string `json:"cloth_name" gorm:"not null"`
	ClothPrice int `json:"cloth_price" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`

	ClothLists []ClothList `json:"cloth_lists" gorm:"foreignKey:ClothID"`
}

func (Cloth) TableName() string {
	return "cloth"
}