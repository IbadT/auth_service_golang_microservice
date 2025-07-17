package auth

import (
	"errors"
	"fmt"

	domain "github.com/IbadT/auth_service_golang_microservice/domain/auth"
	jwtservice "github.com/IbadT/auth_service_golang_microservice/pkg/jwt"
	"github.com/IbadT/auth_service_golang_microservice/pkg/security"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service interface {
	Login(email, password string) (domain.LoginResponse, error)
	Register(email, password string) (domain.User, error)
	RefreshToken(tokenStr string) (domain.LoginResponse, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Login(email, password string) (domain.LoginResponse, error) {
	userExist, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return domain.LoginResponse{}, errors.New("user not found")
	}

	if err := security.ComparePassword(userExist.Password, password); err != nil {
		return domain.LoginResponse{}, err
	}

	tokens, err := jwtservice.CreateToken(userExist.ID, email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return loginResponse, nil
}

func (s *service) Register(email, password string) (domain.User, error) {
	// existingUser, err := s.repo.GetUserByEmail(email)
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		return domain.User{}, errors.New("user is already exist")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		ID:       uuid.New(),
		Email:    email,
		Password: hashedPassword,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

func (s *service) RefreshToken(tokenStr string) (domain.LoginResponse, error) {
	// 1. Проверка токена
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtservice.SecretKey, nil
	})
	if err != nil || !token.Valid {
		return domain.LoginResponse{}, fmt.Errorf("invalid refresh token")
	}

	// 2. Извлечение claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return domain.LoginResponse{}, fmt.Errorf("invalid token claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	email, okEmail := claims["email"].(string)
	if !ok || !okEmail {
		return domain.LoginResponse{}, fmt.Errorf("missing user_id or email in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("invalid user_id format")
	}

	// 3. Генерация новых токенов
	tokens, err := jwtservice.CreateToken(userID, email)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("failed to create new tokens: %w", err)
	}

	return domain.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
