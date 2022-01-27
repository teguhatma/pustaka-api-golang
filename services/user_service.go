package services

import (
	"pustaka-api/repository"
	"pustaka-api/schemas"
)

type UserService interface {
	FindUserById(ID int) (repository.User, error)
	CreateUser(user schemas.UserRequest) (repository.User, error)
	FindUsers() ([]repository.User, error)
	DeleteUser(ID int) (repository.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func UserNewService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindUserById(ID int) (repository.User, error) {
	return s.repository.FindUserById(ID)
}

func (s *userService) CreateUser(userRequest schemas.UserRequest) (repository.User, error) {
	user := repository.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}

func (s *userService) FindUsers() ([]repository.User, error) {
	return s.repository.FindUsers()
}

func (s *userService) DeleteUser(ID int) (repository.User, error) {
	return s.repository.DeleteUser(ID)
}
