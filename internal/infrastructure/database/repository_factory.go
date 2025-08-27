package database

import (
	"database/sql"
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/ports"
	"estudo-go/internal/infrastructure/database/mysql"
	"estudo-go/internal/infrastructure/database/postgres"
	"fmt"
)

func FactoryNewUserRepository(dbType string, db *sql.DB, logger ports.Logger) (domain.UserRepository, error) {
	switch dbType {
	case "mysql":
		repo, err := mysql.NewMySQLUserRepository(db, logger)
		if err != nil {
			return nil, err
		}
		return repo, nil
	case "postgres":
		repo, err := postgres.NewPostgresUserRepository(db, logger)
		if err != nil {
			return nil, err
		}
		return repo, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
