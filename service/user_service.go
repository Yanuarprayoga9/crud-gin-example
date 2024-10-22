package service

import (
	"day1/model/domain"
	"day1/repository"
	"day1/model/web"
	"day1/utils"
	"errors"
)

// UserService interface defines methods for user service
type UserService interface {
	CreateUser(request web.UserCreateRequest) (web.UserResponse, error)
	GetAllUsers() ([]web.UserResponse, error)
	GetUserByID(id uint) (web.UserResponse, error)
	UpdateUser(request web.UserUpdateRequest) (web.UserResponse, error)
	DeleteUser(id uint) error
}

// userService struct implements UserService interface
type userService struct {
	repo      repository.UserRepository
	validator *utils.Validator
}

// NewUserService creates a new instance of userService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo:      repo,
		validator: utils.NewValidator(),
	}
}

// CreateUser validates the request, hashes the password, and calls the repository
func (s *userService) CreateUser(request web.UserCreateRequest) (web.UserResponse, error) {
	// Validate the request
	if err := s.validator.ValidateStruct(request); err != nil {
		return web.UserResponse{}, err
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	// Transform and create the user
	user := &domain.User{
		Username: request.Username,
		Password: hashedPassword, // Store the hashed password
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return web.UserResponse{}, err
	}

	// Prepare the response
	response := web.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
	}

	return response, nil
}

// GetAllUsers retrieves all users and returns them as UserResponse slices
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

// GetUserByID retrieves a single user by ID
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

// UpdateUser validates the request, hashes the password if provided, and calls the repository to update the user
func (s *userService) UpdateUser(request web.UserUpdateRequest) (web.UserResponse, error) {
	// Validate the request
	if err := s.validator.ValidateStruct(request); err != nil {
		return web.UserResponse{}, err
	}

	user, err := s.repo.GetUserByID(uint(request.Id))
	if err != nil {
		return web.UserResponse{}, err
	}

	// Update the user details
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

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	// Check if user exists
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Call the repository to delete the user
	return s.repo.DeleteUser(id)
}
