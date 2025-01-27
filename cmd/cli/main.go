package main

import (
	"clothing-pair-project/internal/config"
	"clothing-pair-project/internal/database/postgres"
	"clothing-pair-project/internal/utils/menu"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadPath()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Failed to close db", err.Error())
		}
	}(db)

	menu.DashboardMenu(db, "")
}
