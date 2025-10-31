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
	err := r.db.Table("cloth").
		Select("id, cloth_name, cloth_price, category").
		Where("category = ?", category).
		Scan(&cloths).Error
	return cloths, err
}

func (r *ClothRepository) GetAllCloths() ([]entity.Cloth, error) {
	var cloths []entity.Cloth
	err := r.db.Table("cloth").
		Select("id, cloth_name, cloth_price, category").
		Scan(&cloths).Error
	return cloths, err
}

func (r *ClothRepository) Update(cloth *entity.Cloth) error {
	return r.db.Table("cloth").
		Where("id = ?", cloth.ID).
		Updates(map[string]interface{}{
			"cloth_name":  cloth.ClothName,
			"cloth_price": cloth.ClothPrice,
			"category":    cloth.Category,
		}).Error
}
