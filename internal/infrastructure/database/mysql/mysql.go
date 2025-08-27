package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("MYSQL_DSN not set")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
