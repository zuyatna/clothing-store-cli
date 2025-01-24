package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

func InitDB(ctx context.Context) (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &DBConfig{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
	}
	if config.Name == "" || config.User == "" || config.Password == "" || config.Host == "" {
		return nil, fmt.Errorf("all database configuration fields are required")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=10",
		config.Host, config.User, config.Password, config.Name)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}
