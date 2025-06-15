package auth

import "gorm.io/gorm"

type Repository interface {
	Login() error
	Register() error
	RefreshToken() error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) Login() error {
	return nil
}

func (r *repository) Register() error {
	return nil
}

func (r *repository) RefreshToken() error {
	return nil
}
