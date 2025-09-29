package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
)

type OrderUsecase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUsecase(orderRepo repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo}
}

func (uc *OrderUsecase) CreateOrder(order *entity.Order) error {
	return uc.orderRepo.CreateOrder(order)
}

func (uc *OrderUsecase) GetOrdersByUserID(userID int) ([]entity.Order, error) {
	return uc.orderRepo.GetByUserID(userID)
}

func (uc *OrderUsecase) GetAllOrders() ([]entity.Order, error) {
	return uc.orderRepo.GetAllOrders()
}

func (uc *OrderUsecase) UpdateOrder(order *entity.Order) error {
	return uc.orderRepo.Update(order)
}