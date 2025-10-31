package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
)

type ClothUsecase struct {
	clothRepo repository.ClothRepository
}

func NewClothUsecase(clothRepo repository.ClothRepository) *ClothUsecase {
	return &ClothUsecase{clothRepo: clothRepo}
}	

func (uc *ClothUsecase) CreateCloth(cloth *entity.Cloth) error {
	return uc.clothRepo.CreateCloth(cloth)
}

func (uc *ClothUsecase) GetClothsByCategory(catagory string) ([]entity.Cloth, error) {
	return uc.clothRepo.GetByCategory(catagory)
}

func (uc *ClothUsecase) GetAllCloths() ([]entity.Cloth, error) {
	return uc.clothRepo.GetAllCloths()
}

func (uc *ClothUsecase) UpdateCloth(cloth *entity.Cloth) error {
	return uc.clothRepo.Update(cloth)
}