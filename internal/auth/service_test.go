package auth

import (
	"errors"
	"testing"

	domain "github.com/IbadT/auth_service_golang_microservice.git/domain/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockRepo) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockRepo) RefreshToken(refresh_token string) error {
	return nil
}

func TestREgister_Success(t *testing.T) {
	repo := new(mockRepo)
	// jwtGen := jwt.CreateToken()
	service := NewService(repo)

	inputEmail := "test@gmail.com"
	inputPassword := "superstongpass"

	// userToCreate := domain.User{
	// 	ID:       uuid.New(), // или uuid.Nil, если ок
	// 	Email:    inputEmail,
	// 	Password: inputPassword,
	// }
	// ожидаем, что пользователя с таким email нет
	repo.On("GetUserByEmail", inputEmail).Return(domain.User{}, errors.New("user not found"))

	// ожидаем успешное создание
	repo.On("CreateUser", mock.AnythingOfType("domain.User")).Return(nil)

	token, err := service.Register(inputEmail, inputPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	repo.AssertExpectations(t)
}
