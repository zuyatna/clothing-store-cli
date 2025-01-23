package main

// go get github.com/olekukonko/tablewriter

import (
	"clothing-pair-project/config"
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	db, err := config.InitDB()

	if err != nil {
		log.Fatal("Error", err.Error())
	}

	defer db.Close()

	sizesMethodHandler := handler.NewSizesHandler(db)
	sizesMethodService := service.NewSizeMethodService(sizesMethodHandler)

	// addSize := entity.Sizes{
	// 	Name: "L",
	// }
	// err = sizesMethodService.Add(addSize)
	// if err != nil {
	// 	log.Fatal("Failed to add size:", err.Error())
	// }
	// log.Println("Successfully added Size:", addSize)

	// updateSize := entity.Sizes{
	// 	Size_id: 1,
	// 	Name:    "XXL",
	// }
	// err = sizesMethodService.Update(updateSize)
	// if err != nil {
	// 	log.Fatal("Failed to update size:", err.Error())
	// }
	// log.Println("Successfully updated Size:", updateSize)

	// idDelete := 1
	// err = sizesMethodService.Delete(idDelete)
	// if err != nil {
	// 	log.Fatal("Failed to delete Size:", err.Error())
	// }
	// log.Println("Successfully deleted Size ID", idDelete)

	// sizeFindAll, err := sizesMethodService.FindAll()
	// if err != nil {
	// 	log.Fatal("Failed to find all size:", err.Error())
	// }

	sizeId := 1
	sizeFindOne, err := sizesMethodService.Find(&sizeId)
	if err != nil {
		log.Fatal("Failed to find data : ", err.Error())
	}

	ShowData("Size", sizeFindOne)

}

func ShowData(namatable string, sizes []entity.Sizes) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("ID\tName\n"))
	_, _ = w.Write([]byte("--\t----\n"))

	for _, size := range sizes {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%d\t%s\n", size.Size_id, size.Name),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
}
