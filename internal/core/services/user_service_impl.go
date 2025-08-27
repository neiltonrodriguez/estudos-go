package services

import (
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/ports"
	"fmt"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepo domain.UserRepository
	logger   ports.Logger
}

func NewUserService(userRepo domain.UserRepository, logger ports.Logger) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *userServiceImpl) CreateUser(name, email, password string) (*domain.User, error) {
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil && err != domain.ErrUserNotFound {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	id := uuid.New().String()

	user, err := domain.NewUser(id, name, email, password)
	if err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	if err := s.userRepo.Save(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}
	s.logger.Info("userService: create user with successfully")

	return user, nil
}

func (s *userServiceImpl) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	s.logger.Info("userService: get user: ", user)
	return user, nil
}
