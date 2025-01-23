package main

// go get github.com/olekukonko/tablewriter

import (
	"clothing-pair-project/config"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"log"
)

func main() {
	db, err := config.InitDB()

	if err != nil {
		log.Fatal("Error", err.Error())
	}

	defer db.Close()

	catMethodHandler := handler.NewCategoryHandler(db)
	catMethodService := service.NewCategoryMethodService(catMethodHandler)

	// addCategory := entity.Categories{
	// 	Collection_id: 1,
	// 	Name:          "Pants",
	// }
	// err = catMethodService.Add(addCategory)
	// if err != nil {
	// 	log.Fatal("Failed to add:", err.Error())
	// }
	// log.Println("Successfully added:", addCategory)

	// updateCategory := entity.Categories{
	// 	Category_id:   3,
	// 	Collection_id: 1,
	// 	Name:          "Pant",
	// }
	// err = catMethodService.Update(updateCategory)
	// if err != nil {
	// 	log.Fatal("Failed to update:", err.Error())
	// }
	// log.Println("Successfully updated:", updateCategory)

	// idDelete := 3
	// err = catMethodService.Delete(idDelete)
	// if err != nil {
	// 	log.Fatal("Failed to delete :", err.Error())
	// }
	// log.Println("Successfully deleted ", idDelete)

	sizeId := 3
	sizeFind, err := catMethodService.Find(&sizeId)
	if err != nil {
		log.Fatal("Failed to find data : ", err.Error())
	}

	handler.ShowDataCategory("Category", sizeFind)

}
