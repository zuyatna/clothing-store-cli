package main

import (
	"clothing-pair-project/internal/config"
	"clothing-pair-project/internal/database/postgres"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	fmt.Println("Successfully connected to database")
}
