package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/messages"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageCategoryMenu(db *sqlx.DB, message string) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Manage Category Menu")
	fmt.Println("1. Find All Categories")
	fmt.Println("2. Add Category")
	fmt.Println("3. Edit Category")
	fmt.Println("4. Delete Category")
	fmt.Println("0. Back")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageCategoryMenu(db, message)
	}

	categoryRepository := sql.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)

	switch input {
	case "1":
		findAllCategoriesMenu(db, categoryService, "")
	case "2":
		addCategoryMenu(db, categoryService, "")
	case "3":
		editCategoryMenu(db, categoryService, "")
	case "4":
		deleteCategoryMenu(db, categoryService, "")
	case "0":
		AdminMenu(db, "")
	default:
		message = "Invalid input"
		messages.PrintMessage(message)
		ManageCategoryMenu(db, message)
	}
}

func allCategories(categoryService *services.CategoryService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})
	table.SetRowLine(true)

	categories, err := categoryService.GetAllCategories()
	if err != nil {
		fmt.Println("Failed to get all categories")
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories found")
		return
	}

	for _, category := range categories {
		table.Append([]string{strconv.Itoa(category.CategoryID), category.Name})
	}
	table.Render()
	fmt.Println()
}

func findAllCategoriesMenu(db *sqlx.DB, categoryService *services.CategoryService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Find All Categories")
	fmt.Println("=====================================")

	allCategories(categoryService)

	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		ManageCategoryMenu(db, message)
	}

	ManageCategoryMenu(db, "")
}

func addCategoryMenu(db *sqlx.DB, categoryService *services.CategoryService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Add Category")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter category name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	category := models.Category{
		Name: name,
	}

	err := categoryService.AddCategory(category)
	if err != nil {
		fmt.Printf("Error adding category: %v\n", err)
		return
	}

	ManageCategoryMenu(db, "Category added successfully")
}

func editCategoryMenu(db *sqlx.DB, categoryService *services.CategoryService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Edit Category")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	reader := bufio.NewReader(os.Stdin)

	allCategories(categoryService)

	fmt.Print("Enter category ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid category ID")
		editCategoryMenu(db, categoryService, message)
	}

	category, err := categoryService.GetCategoryByID(id)
	if err != nil {
		fmt.Printf("Error finding category: %v\n", err)
		return
	}

	fmt.Printf("Current name: %s\n", category.Name)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	category.Name = name

	err = categoryService.UpdateCategory(category)
	if err != nil {
		fmt.Printf("Error updating category: %v\n", err)
		return
	}

	ManageCategoryMenu(db, "Category updated successfully")
}

func deleteCategoryMenu(db *sqlx.DB, categoryService *services.CategoryService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Delete Category")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	reader := bufio.NewReader(os.Stdin)

	allCategories(categoryService)

	fmt.Print("Enter category ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid category ID")
		deleteCategoryMenu(db, categoryService, message)
	}

	err = categoryService.DeleteCategory(id)
	if err != nil {
		fmt.Printf("Error deleting category: %v\n", err)
		return
	}

	ManageCategoryMenu(db, "Category deleted successfully")
}
