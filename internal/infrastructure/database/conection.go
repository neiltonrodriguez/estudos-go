package database

import (
	"database/sql"
	"fmt"
	"os"

	"estudo-go/internal/infrastructure/database/mysql"
	"estudo-go/internal/infrastructure/database/postgres"
)

func NewConnection() (*sql.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "mysql":
		return mysql.NewMySQLConnection()
	case "postgres":
		return postgres.NewPostgresConnection()
	default:
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", dbType)
	}
}
