package services_test

import (
	"errors"
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository é uma implementação mock da interface UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	t.Run("should create a user successfully", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		// Esperamos que GetByEmail seja chamado e retorne nil (usuário não existe)
		mockRepo.On("GetByEmail", "test@example.com").Return(nil, nil).Once()
		// Esperamos que Save seja chamado e retorne nil (sucesso ao salvar)
		mockRepo.On("Save", mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userService := services.NewUserService(mockRepo) // Assume que teremos um NewUserService

		user, err := userService.CreateUser("Test User", "test@example.com", "password123")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)
		assert.Equal(t, "test@example.com", user.Email)
		mockRepo.AssertExpectations(t) // Verifica se os métodos mockados foram chamados
	})

	t.Run("should return error if user with email already exists", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		existingUser := &domain.User{Email: "existing@example.com"}
		// Esperamos que GetByEmail seja chamado e retorne um usuário existente
		mockRepo.On("GetByEmail", "existing@example.com").Return(existingUser, nil).Once()

		userService := services.NewUserService(mockRepo)

		user, err := userService.CreateUser("Another User", "existing@example.com", "password123")

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, services.ErrUserAlreadyExists, err) // Assume que definiremos este erro
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if Save fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		// GetByEmail retorna nil
		mockRepo.On("GetByEmail", "fail@example.com").Return(nil, nil).Once()
		// Save retorna um erro
		mockRepo.On("Save", mock.AnythingOfType("*domain.User")).Return(errors.New("db error")).Once()

		userService := services.NewUserService(mockRepo)

		user, err := userService.CreateUser("Fail User", "fail@example.com", "password123")

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "db error")
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error for invalid user fields", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userService := services.NewUserService(mockRepo) // Não precisa mockar GetByEmail ou Save

		user, err := userService.CreateUser("", "invalid@example.com", "password123") // Nome vazio

		assert.Nil(t, user)
		assert.NotNil(t, err)
		// assert.IsType(t, domain.ErrInvalidUserFields{}, err)     // Verifica o tipo de erro
		mockRepo.AssertNotCalled(t, "GetByEmail", mock.Anything) // Não deve chamar GetByEmail
		mockRepo.AssertNotCalled(t, "Save", mock.Anything)       // Não deve chamar Save
	})
}
