package repository

import (
	"day1/model/domain"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *domain.User) error
    GetAllUsers() ([]domain.User, error)
    GetUserByID(id uint) (domain.User, error)
    UpdateUser(user *domain.User) error
    DeleteUser(id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]domain.User, error) {
    var users []domain.User
    err := r.db.Find(&users).Error
    for _, v := range users {
        fmt.Println(v.Username)
        fmt.Println(v.Password)
    }
    return users, err
}

func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return user, err
}

func (r *userRepository) UpdateUser(user *domain.User) error {
    return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
    return r.db.Delete(&domain.User{}, id).Error
}
