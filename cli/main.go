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

	productMethodHandler := handler.NewProductsHandler(db)
	ProductMethodService := service.NewProductMethodService(productMethodHandler)

	// addProduct := entity.Products{
	// 	Category_Id: 1,
	// 	Color_Id:    1,
	// 	Size_Id:     2,
	// 	Price:       20000,
	// 	Stock:       20,
	// 	Description: "Baju",
	// 	Image:       "/asset/product/asd.jpg",
	// }
	// err = ProductMethodService.Add(addProduct)
	// if err != nil {
	// 	log.Fatal("Failed to add:", err.Error())
	// }
	// log.Println("Successfully added:", addProduct)

	// updateProduct := entity.Products{
	// 	Product_Id:  3,
	// 	Category_Id: 1,
	// 	Color_Id:    1,
	// 	Size_Id:     2,
	// 	Price:       20000,
	// 	Stock:       10,
	// 	Description: "ju",
	// 	Image:       "/asset/product/asd.jpg",
	// }
	// err = ProductMethodService.Update(updateProduct)
	// if err != nil {
	// 	log.Fatal("Failed to update:", err.Error())
	// }
	// log.Println("Successfully updated:", updateProduct)

	// idDelete := 3
	// err = ProductMethodService.Delete(idDelete)
	// if err != nil {
	// 	log.Fatal("Failed to delete :", err.Error())
	// }
	// log.Println("Successfully deleted ", idDelete)

	// sizeId := 3
	sizeFind, err := ProductMethodService.Find(nil)
	if err != nil {
		log.Fatal("Failed to find data : ", err.Error())
	}

	handler.ShowDataProduct("Product", sizeFind)

}
