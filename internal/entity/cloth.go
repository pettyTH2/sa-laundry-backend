package entity

import "gorm.io/gorm"

type Cloth struct {
	gorm.Model
	ClothName string `json:"cloth_name"`
	ClothPrice int `json:"cloth_price"`
	Catagory string `json:"category"`

	ClothLists []ClothList `json:"cloth_lists" gorm:"foreignKey:ClothID"`
}