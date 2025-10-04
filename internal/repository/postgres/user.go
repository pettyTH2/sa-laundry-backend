package postgres

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return r.db.Create(user).Error
}

func (r *UserRepository) UserLogin(phoneNumber, password string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("phone_number = ?", phoneNumber).First(&user)
	
	if result.Error != nil {
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	
	return &user, nil
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