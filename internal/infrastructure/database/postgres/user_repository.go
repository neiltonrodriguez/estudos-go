package postgres

import (
	"database/sql"
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/ports"
	"fmt"
)

type PostgresUserRepository struct {
	db     *sql.DB
	logger ports.Logger
}

func NewPostgresUserRepository(db *sql.DB, logger ports.Logger) (*PostgresUserRepository, error) {
	return &PostgresUserRepository{db: db, logger: logger}, nil
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		r.logger.Error("failed to save user: %v", err)
		return fmt.Errorf("failed to save user to Postgres: %w", err)
	}
	r.logger.Info("user saved: %s", user.Email)
	return nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		r.logger.Error("failed to get user by email: %v", err)
		return nil, fmt.Errorf("failed to get user by email from Postgres: %w", err)
	}
	r.logger.Info("Get user with successfully", user.Email)
	return user, nil
}

func (r *PostgresUserRepository) Close() error {
	return r.db.Close()
}
