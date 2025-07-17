package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" validata:"uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=6"`
}

func NewUserUUID(uid uuid.UUID) *User {
	return &User{
		ID: uid,
	}
}

func NewUser(email, password string) *User {

	return &User{
		Email:    email,
		Password: password,
	}
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    time.Time
}

func NewLoginResponse(access_token, refresh_token string) *LoginResponse {
	return &LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
}
