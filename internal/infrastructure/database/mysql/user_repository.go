package mysql

import (
	"database/sql"
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/ports"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLUserRepository struct {
	db     *sql.DB
	logger ports.Logger
}

func NewMySQLUserRepository(db *sql.DB, logger ports.Logger) (*MySQLUserRepository, error) {
	return &MySQLUserRepository{db: db, logger: logger}, nil
}

func (r *MySQLUserRepository) Save(user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save user to database: %w", err)
	}
	r.logger.Info("User %s saved successfully", user.Email)
	return nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email from database: %w", err)
	}
	r.logger.Info("Get user with successfully", user.Email)
	return user, nil
}

func (r *MySQLUserRepository) Close() error {
	return r.db.Close()
}
