package menu

import (
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ManageProductMenu(db *sqlx.DB, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Manage Product Menu")
	fmt.Println("1. Find All Products")
	fmt.Println("2. Find Product By Name")
	fmt.Println("3. Find Product By Category")
	fmt.Println("4. Add Product")
	fmt.Println("5. Update Product")
	fmt.Println("6. Delete Product")
	fmt.Println("0. Back")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageUserMenu(db, message)
	}

	// productRepository := sql.NewProductRepository(db)
	// productService := services.NewProductService(productRepository)

	switch input {
	case "1":
		// Find All Products
	case "2":
		// Find Product By Name
	case "3":
		// Find Product By Category
	case "4":
		// Add Product
	case "5":
		// Update Product
	case "6":
	// Delete Product
	default:
		message = "Invalid input"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}
}
