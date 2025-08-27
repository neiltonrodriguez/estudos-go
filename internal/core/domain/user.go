package domain

import (
	"errors"
	"time"
)

// User representa a entidade de usuário no domínio
type User struct {
	ID        string    `json:"id"` // Usaremos UUID para ID
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Oculta a senha ao serializar para JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser cria uma nova instância de User com validações básicas
func NewUser(id, name, email, password string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDRequired
	}
	if name == "" {
		return nil, ErrUserNameRequired
	}
	if email == "" {
		return nil, ErrUserEmailRequired
	}
	if password == "" {
		return nil, ErrUserPasswordRequired
	}

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Erros customizados para o domínio
var (
	ErrUserIDRequired       = errors.New("user ID is required")
	ErrUserNameRequired     = errors.New("user name is required")
	ErrUserEmailRequired    = errors.New("user email is required")
	ErrUserPasswordRequired = errors.New("user password is required")
	ErrInvalidUserFields    = errors.New("invalid user fields")
	ErrUserNotFound         = errors.New("user not found")
)
