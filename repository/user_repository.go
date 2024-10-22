package repository

import (
	"day1/model/domain"
	"fmt"

	"gorm.io/gorm"
)

// UserRepository interface defines methods for user repository
type UserRepository interface {
    CreateUser(user *domain.User) error
    GetAllUsers() ([]domain.User, error)
    GetUserByID(id uint) (domain.User, error)
    UpdateUser(user *domain.User) error
    DeleteUser(id uint) error
}

// userRepository struct implements UserRepository interface
type userRepository struct {
    db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *userRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

// GetAllUsers retrieves all users from the database
func (r *userRepository) GetAllUsers() ([]domain.User, error) {
    var users []domain.User
    err := r.db.Find(&users).Error
    for _, v := range users {
        fmt.Println(v.Username)
        fmt.Println(v.Password)
    }
    return users, err
}

// GetUserByID retrieves a user by ID from the database
func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return user, err
}

// UpdateUser updates a user in the database
func (r *userRepository) UpdateUser(user *domain.User) error {
    return r.db.Save(user).Error
}

// DeleteUser deletes a user by ID from the database
func (r *userRepository) DeleteUser(id uint) error {
    return r.db.Delete(&domain.User{}, id).Error
}
