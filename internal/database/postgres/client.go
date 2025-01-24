package postgres

import (
	"clothing-pair-project/internal/config"

	"github.com/jmoiron/sqlx"
)

func InitDB(cfg *config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	return db, err
}
