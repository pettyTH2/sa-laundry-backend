package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ClothListRepository struct {	
	db *gorm.DB
}

func NewClothListRepository(db *gorm.DB) repository.ClothListRepository {
	return &ClothListRepository{db: db}
}

func (r *ClothListRepository) CreateClothList(clothList *entity.ClothList) error {	
	return r.db.Create(clothList).Error
}

func (r *ClothListRepository) GetByOrderID(orderID uuid.UUID) ([]entity.ClothList, error) {
	var clothLists []entity.ClothList
	err := r.db.Preload("Order").Preload("Cloth").Where("order_id = ?", orderID).Find(&clothLists).Error
	return clothLists, err
}

func (r *ClothListRepository) Update(clothList *entity.ClothList) error {
	return r.db.Save(clothList).Error
}