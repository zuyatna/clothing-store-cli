package menu

import (
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageSizeMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage Size Menu")
		fmt.Println("1. Add Size")
		fmt.Println("2. Find All Sizes")
		fmt.Println("3. Find Size By Name")
		fmt.Println("4. Update Size")
		fmt.Println("5. Delete Size")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		sizeHandler := handler.NewSizesHandler(db)
		sizeService := service.NewSizeMethodService(sizeHandler)

		switch input {
		case 1:
			// TODO: Add size menu
		case 2:
			findAllSizes(sizeService)
		case 3:
			// TODO: Find size by name menu
		case 4:
			// TODO: Update size menu
		case 5:
			// TODO: Delete size menu
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func findAllSizes(sizeService *service.SizeMethodService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Sizes")
	fmt.Println("=====================================")

	AllSizes(sizeService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func AllSizes(sizeService *service.SizeMethodService) {
	sizes, err := sizeService.Find(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(sizes) == 0 {
		fmt.Println("no data found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Size ID", "Name"})
	for _, size := range sizes {
		table.Append([]string{fmt.Sprint(size.Size_id), size.Name})
	}
	table.Render()

	fmt.Println()
}
