package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
