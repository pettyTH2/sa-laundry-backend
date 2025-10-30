package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByUserID(userID int) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Preload("User").Preload("ClothList").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Preload("User").Preload("ClothList").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) Update(order *entity.Order) error {
	return r.db.Save(order).Error
}