package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string `json:"host" toml:"host"`
	Port     int    `json:"port" toml:"port"`
	User     string `json:"user" toml:"user"`
	Password string `json:"password" toml:"password"`
	DBName   string `json:"dbname" toml:"dbname"`
	SSLMode  string `json:"sslmode" toml:"sslmode"`
}

func Initialize(cfg *DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
