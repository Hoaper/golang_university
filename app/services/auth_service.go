package services

import (
	"errors"
	"github.com/Hoaper/golang_university/app/models"
	"github.com/Hoaper/golang_university/app/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	AuthenticateUser(login, password string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepository: *userRepository}
}

func (s *AuthService) GetUserByLogin(login string) (*models.User, error) {
	user, err := s.UserRepository.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) CreateUser(user *models.User) error {
	findUser, _ := s.UserRepository.GetUserByLogin(user.Login)

	if findUser != nil {
		return errors.New("user already exists")
	}
	return s.UserRepository.CreateUser(user)
}

func (s *AuthService) AuthenticateUser(login, password string) (*models.User, error) {
	user, err := s.UserRepository.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, err
	}

	return user, nil
}
