package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"gorm.io/gorm"
)

type ClothRepository struct {
	db *gorm.DB
}

func NewClothRepository(db *gorm.DB) repository.ClothRepository {
	return &ClothRepository{db: db}
}

func (r *ClothRepository) CreateCloth(cloth *entity.Cloth) error {
	return r.db.Create(cloth).Error
}

func (r *ClothRepository) GetByCategory(category string) ([]entity.Cloth, error) {
	var cloths []entity.Cloth
	err := r.db.Where("category = ?", category).Find(&cloths).Error
	return cloths, err
}

func (r *ClothRepository) GetAllCloths() ([]entity.Cloth, error) {
	var cloths []entity.Cloth
	err := r.db.Find(&cloths).Error
	return cloths, err
}

func (r *ClothRepository) Update(cloth *entity.Cloth) error {
	return r.db.Save(cloth).Error
}