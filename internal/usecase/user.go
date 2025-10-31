package usecase

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) CreateUser(user *entity.User) error {
	return uc.userRepo.CreateUser(user)
}

func (uc *UserUsecase) UserLogin(phoneNumber, password string) (*entity.User, error) {
	return uc.userRepo.UserLogin(phoneNumber, password)
}

func (uc *UserUsecase) GetUserByRole(role string) ([]entity.User, error) {
	return uc.userRepo.GetByRole(role)
}

func (uc *UserUsecase) GetUserById(id string) (*entity.User, error) {
	return uc.userRepo.GetById(id)
}

func (uc *UserUsecase) GetAllUsers() ([]entity.User, error) {
	return uc.userRepo.GetAllUsers()
}

func (uc *UserUsecase) UpdateUser(user *entity.User) error {
	return uc.userRepo.Update(user)
}

