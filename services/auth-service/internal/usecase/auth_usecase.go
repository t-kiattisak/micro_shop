package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginResult struct {
	Token string
	User  *domain.User
}

type AuthUsecase interface {
	Register(username, password string) (*domain.User, error)
	Login(username, password string) (*LoginResult, error)
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

func (a *authUsecase) Login(username, password string) (*LoginResult, error) {
	user, err := a.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Token: tokenStr,
		User:  user,
	}, nil
}
