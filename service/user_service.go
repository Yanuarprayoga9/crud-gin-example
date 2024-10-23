package service

import (
	"day1/model/domain"
	"day1/repository"
	"day1/model/web"
	"day1/utils"
	"errors"
)

type UserService interface {
	CreateUser(request web.UserCreateRequest) (web.UserResponse, error)
	GetAllUsers() ([]web.UserResponse, error)
	GetUserByID(id uint) (web.UserResponse, error)
	UpdateUser(request web.UserUpdateRequest) (web.UserResponse, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo      repository.UserRepository
	validator *utils.Validator
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo:      repo,
		validator: utils.NewValidator(),
	}
}

func (s *userService) CreateUser(request web.UserCreateRequest) (web.UserResponse, error) {
	if err := s.validator.ValidateStruct(request); err != nil {
		return web.UserResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := &domain.User{
		Username: request.Username,
		Password: hashedPassword, // Store the hashed password
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return web.UserResponse{}, err
	}

	response := web.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
	}

	return response, nil
}

func (s *userService) GetAllUsers() ([]web.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []web.UserResponse
	for _, user := range users {
		responses = append(responses, web.UserResponse{
			Id:       int(user.ID),
			Username: user.Username,
		})
	}

	return responses, nil
}

func (s *userService) GetUserByID(id uint) (web.UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
	}, nil
}

func (s *userService) UpdateUser(request web.UserUpdateRequest) (web.UserResponse, error) {
	if err := s.validator.ValidateStruct(request); err != nil {
		return web.UserResponse{}, err
	}

	user, err := s.repo.GetUserByID(uint(request.Id))
	if err != nil {
		return web.UserResponse{}, err
	}

	user.Username = request.Username
	if request.Password != "" {
		hashedPassword, err := utils.HashPassword(request.Password) // Hash the password if provided
		if err != nil {
			return web.UserResponse{}, err
		}
		user.Password = hashedPassword
	}

	err = s.repo.UpdateUser(&user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
	}, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.repo.DeleteUser(id)
}
