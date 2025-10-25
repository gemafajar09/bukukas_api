package usecase

import (
	"errors"
	"fmt"
	"go-project/internal/auth"
	"go-project/internal/domain"
	"go-project/internal/repository"
)

type UserUsecase interface {
	Register(user domain.User) (domain.User, error)
	Login(email, password string) (string, error)
	GetUserByID(i uint) (domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) Register(user domain.User) (domain.User, error) {
	user.Password = auth.HashPassword(user.Password)
	return uc.userRepository.CreateUser(user)
}

func (uc *userUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUsecase) GetUserByID(i uint) (domain.User, error) {
	user, err := uc.userRepository.GetUserByID(i)
	fmt.Println(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}

func (uc *userUsecase) GetUsers() ([]domain.User, error) {
	users, err := uc.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, err
}
