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

func (r *UserRepository) UserLogin(Id, password string) (*entity.User, error) {
	var user entity.User
	result := r.db.Table("user").
		Select("id, phone_number, name, nickname, password, role").
		Where("phone_number = ?", Id).
		Limit(1).
		Scan(&user)

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
	err := r.db.Table("user").
		Select("id, phone_number, name, nickname, password, role").
		Where("role = ?", Role).
		Scan(&users).Error
	return users, err
}

func (r *UserRepository) GetById(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("user").
		Select("id, phone_number, name, nickname, password, role").
		Where("id = ?", id).
		Scan(&user).Error
	return &user, err
}

func (r *UserRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Table("user").
		Select("id, phone_number, name, nickname, password, role").
		Scan(&users).Error
	return users, err
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.db.Table("user").
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"phone_number": user.PhoneNumber,
			"name":         user.Name,
			"nickname":     user.Nickname,
			"password":     user.Password,
			"role":         user.Role,
		}).Error
}
