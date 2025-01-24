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

func ManageCategoryMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage Category Menu")
		fmt.Println("1. Add Category")
		fmt.Println("2. Find All Categories")
		fmt.Println("3. Find Category By Name")
		fmt.Println("4. Update Category")
		fmt.Println("5. Delete Category")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		categoryHandler := handler.NewCategoryHandler(db)
		categoryService := service.NewCategoryMethodService(categoryHandler)

		switch input {
		case 1:
			// TODO: Add category menu
		case 2:
			findAllCategories(categoryService)
		case 3:
			// TODO: Find category by name menu
		case 4:
			// TODO: Update category menu
		case 5:
			// TODO: Delete category menu
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func findAllCategories(categoryService *service.CategoryMethodService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Categories")
	fmt.Println("=====================================")

	AllCategory(categoryService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func AllCategory(categoryService *service.CategoryMethodService) {
	categories, err := categoryService.Find(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})
	for _, category := range categories {
		table.Append([]string{strconv.Itoa(category.Category_id), category.Name})
	}
	table.Render()

	fmt.Println()
}
