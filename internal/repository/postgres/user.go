package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByRole(Role string) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Where("role = ?", Role).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone_number = ?", phoneNumber).Find(&user).Error
	return &user, err
}

func (r *UserRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}