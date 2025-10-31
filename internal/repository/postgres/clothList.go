package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	err := r.db.Table("cloth_list").
		Select("id, quantity, sub_total_cost, order_id, cloth_id").
		Where("order_id = ?", orderID).
		Scan(&clothLists).Error

	if err != nil {
		return clothLists, err
	}

	// Load Cloth for each ClothList
	for i := range clothLists {
		r.db.Table("cloth").
			Select("id, cloth_name, cloth_price, category").
			Where("id = ?", clothLists[i].ClothID).
			Scan(&clothLists[i].Cloth)
	}

	return clothLists, nil
}

func (r *ClothListRepository) Update(clothList *entity.ClothList) error {
	return r.db.Table("cloth_list").
		Where("id = ?", clothList.ID).
		Updates(map[string]interface{}{
			"quantity":       clothList.Quantity,
			"sub_total_cost": clothList.SubTotalCost,
			"order_id":       clothList.OrderID,
			"cloth_id":       clothList.ClothID,
		}).Error
}
