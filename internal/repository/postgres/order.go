package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"

	"github.com/google/uuid"
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

func (r *OrderRepository) GetByUserID(userID uuid.UUID) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Table("order").
		Select("id, service_type, total_cloth, total_cost, order_date, pickup_date, order_status, payment_method, user_id").
		Where("user_id = ?", userID).
		Scan(&orders).Error

	if err != nil {
		return orders, err
	}

	// Load User for each order
	for i := range orders {
		r.db.Table("user").
			Select("id, phone_number, name, nickname, role").
			Where("id = ?", orders[i].UserID).
			Scan(&orders[i].User)

		// Load ClothLists
		var clothLists []entity.ClothList
		r.db.Table("cloth_list").
			Select("id, quantity, sub_total_cost, order_id, cloth_id").
			Where("order_id = ?", orders[i].ID).
			Scan(&clothLists)

		// Load Cloth for each ClothList
		for j := range clothLists {
			r.db.Table("cloth").
				Select("id, cloth_name, cloth_price, category").
				Where("id = ?", clothLists[j].ClothID).
				Scan(&clothLists[j].Cloth)
		}

		orders[i].ClothLists = clothLists
	}

	return orders, nil
}

func (r *OrderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Table("order").
		Select("id, service_type, total_cloth, total_cost, order_date, pickup_date, order_status, payment_method, user_id").
		Scan(&orders).Error

	if err != nil {
		return orders, err
	}

	// Load User for each order
	for i := range orders {
		r.db.Table("user").
			Select("id, phone_number, name, nickname, role").
			Where("id = ?", orders[i].UserID).
			Scan(&orders[i].User)

		// Load ClothLists
		var clothLists []entity.ClothList
		r.db.Table("cloth_list").
			Select("id, quantity, sub_total_cost, order_id, cloth_id").
			Where("order_id = ?", orders[i].ID).
			Scan(&clothLists)

		// Load Cloth for each ClothList
		for j := range clothLists {
			r.db.Table("cloth").
				Select("id, cloth_name, cloth_price, category").
				Where("id = ?", clothLists[j].ClothID).
				Scan(&clothLists[j].Cloth)
		}

		orders[i].ClothLists = clothLists
	}

	return orders, nil
}

func (r *OrderRepository) Update(order *entity.Order) error {
	return r.db.Table("order").
		Where("id = ?", order.ID).
		Updates(map[string]interface{}{
			"service_type":   order.ServiceType,
			"total_cloth":    order.TotalCloth,
			"total_cost":     order.TotalCost,
			"order_date":     order.OrderDate,
			"pickup_date":    order.PickupDate,
			"order_status":   order.OrderStatus,
			"payment_method": order.PaymentMethod,
			"user_id":        order.UserID,
		}).Error
}

func (r *OrderRepository) GetByID(id uuid.UUID) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Table("order").
		Select("id, service_type, total_cloth, total_cost, order_date, pickup_date, order_status, payment_method, user_id").
		Where("id = ?", id).
		Limit(1).
		Scan(&order).Error

	if err != nil {
		return nil, err
	}

	// Check if order exists
	if order.ID == uuid.Nil {
		return nil, gorm.ErrRecordNotFound
	}

	// Load User
	r.db.Table("user").
		Select("id, phone_number, name, nickname, role").
		Where("id = ?", order.UserID).
		Scan(&order.User)

	// Load ClothLists
	var clothLists []entity.ClothList
	r.db.Table("cloth_list").
		Select("id, quantity, sub_total_cost, order_id, cloth_id").
		Where("order_id = ?", order.ID).
		Scan(&clothLists)

	// Load Cloth for each ClothList
	for i := range clothLists {
		r.db.Table("cloth").
			Select("id, cloth_name, cloth_price, category").
			Where("id = ?", clothLists[i].ClothID).
			Scan(&clothLists[i].Cloth)
	}

	order.ClothLists = clothLists

	return &order, nil
}
