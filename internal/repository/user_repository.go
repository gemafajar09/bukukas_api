package repository

import (
	"go-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUserByID(id uint) (domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
