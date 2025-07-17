package auth

import (
	domain "github.com/IbadT/auth_service_golang_microservice.git/domain/auth"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (domain.User, error)
	CreateUser(user domain.User) error
	RefreshToken(refresh_token string) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.DB.First(&user, "email = ?", email).Error
	return user, err
}

func (r *repository) CreateUser(user domain.User) error {
	return r.DB.Create(user).Error
}

func (r *repository) RefreshToken(refresh_token string) error {
	return nil
}
