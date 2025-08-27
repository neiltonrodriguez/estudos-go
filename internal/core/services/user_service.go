package services

import (
	"errors"
	"estudo-go/internal/core/domain"
)

// UserService define a interface para o serviço de usuário.
// Esta camada contém a lógica de negócio principal.
type UserService interface {
	CreateUser(name, email, password string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

// Erros customizados para o serviço
var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
)
