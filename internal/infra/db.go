package infra

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

// InitDB отвечает за подключение к БД и возвращает `*sql.DB`.
func InitDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DSN is not set in environment variables")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}
