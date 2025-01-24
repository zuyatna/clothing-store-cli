package main

import (
	"clothing-pair-project/internal/config"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	db, err := config.InitDB(ctx)
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()
}
