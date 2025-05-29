package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(username, password string) (*domain.User, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (a *authUsecase) Register(username string, password string) (*domain.User, error) {
	existing, _ := a.repo.FindByUsername(username)
	if existing != nil {
		return nil, errors.New("username already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         "user",
	}

	if err := a.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
