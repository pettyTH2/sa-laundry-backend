package repository

import "laundry-backend/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetByRole(role string) ([]entity.User, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	Update(user *entity.User) error
}