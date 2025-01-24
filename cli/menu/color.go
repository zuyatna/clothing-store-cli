package menu

import (
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageColorMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage Color Menu")
		fmt.Println("1. Add Color")
		fmt.Println("2. Find All Colors")
		fmt.Println("3. Find Color By Name")
		fmt.Println("4. Update Color")
		fmt.Println("5. Delete Color")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		colorHandler := handler.NewColorHandler(db)
		colorService := service.NewColorService(colorHandler)

		switch input {
		case 1:
			// TODO: Add color menu
		case 2:
			findAllColors(colorService)
		case 3:
			// TODO: Find color by name menu
		case 4:
			// TODO: Update color menu
		case 5:
			// TODO: Delete color menu
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func findAllColors(colorService *service.ColorService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Colors")
	fmt.Println("=====================================")

	AllColor(colorService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func AllColor(colorService *service.ColorService) {
	colors, err := colorService.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	if len(colors) == 0 {
		fmt.Println("No colors found")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})
	for _, color := range colors {
		table.Append([]string{strconv.Itoa(color.ColorID), color.Name})
	}
	table.Render()

	fmt.Println()
}
