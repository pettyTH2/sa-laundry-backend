package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"github.com/google/uuid"
)

type ClothListUsecase struct {
	clothListRepo repository.ClothListRepository
}

func NewClothListUsecase(clothListRepo repository.ClothListRepository) *ClothListUsecase {
	return &ClothListUsecase{clothListRepo: clothListRepo}
}

func (uc *ClothListUsecase) CreateClothList(clothList *entity.ClothList) error {
	return uc.clothListRepo.CreateClothList(clothList)
}

func (uc *ClothListUsecase) GetClothListsByOrderID(orderID uuid.UUID) ([]entity.ClothList, error) {
	return uc.clothListRepo.GetByOrderID(orderID)
}

func (uc *ClothListUsecase) UpdateClothList(clothList *entity.ClothList) error {
	return uc.clothListRepo.Update(clothList)
}