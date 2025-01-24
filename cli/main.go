package main

import (
	"clothing-pair-project/config"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"log"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	productMethodHandler := handler.NewProductsHandler(db)
	productMethodService := service.NewProductMethodService(productMethodHandler)

	sizeId := 1
	sizeFindOne, err := productMethodService.Find(&sizeId)
	if err != nil {
		log.Fatal("Failed to find data : ", err.Error())
	}
	handler.ShowDataProduct("Products", sizeFindOne)
}
