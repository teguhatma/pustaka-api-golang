package services

import (
	"pustaka-api/repository"
	"pustaka-api/schemas"
	"pustaka-api/utils"
)

type UserService interface {
	FindUserById(ID int) (schemas.User, error)
	CreateUser(user schemas.UserRequest) (repository.User, error)
	FindUsers() ([]schemas.User, error)
	DeleteUser(ID int) (repository.User, error)
	UpdateUser(user schemas.UserRequest, ID int) (repository.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func UserNewService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindUserById(ID int) (schemas.User, error) {
	return s.repository.FindUserById(ID)
}

func (s *userService) CreateUser(userRequest schemas.UserRequest) (repository.User, error) {
	password, _ := utils.HashPassword(userRequest.Password)
	user := repository.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: password,
	}
	res, err := s.repository.CreateUser(user)
	return res, err
}

func (s *userService) FindUsers() ([]schemas.User, error) {
	// var users []schemas.UserResponse

	res, err := s.repository.FindUsers()

	return res, err
}

func (s *userService) DeleteUser(ID int) (repository.User, error) {
	return s.repository.DeleteUser(ID)
}

func (s *userService) UpdateUser(userRequest schemas.UserRequest, ID int) (repository.User, error) {
	password, _ := utils.HashPassword(userRequest.Password)
	user := repository.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: password,
	}

	res, err := s.repository.UpdateUser(user, ID)
	return res, err
}
