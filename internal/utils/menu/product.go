package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/terminal"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
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

	productRepository := sql.NewProductRepository(db)
	productService := services.NewProductService(productRepository)

	switch input {
	case "1":
		findAllProductsMenu(db, productService, "")
	case "2":
		findProductByNameMenu(db, productService, "")
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

func allProducts(productService *services.ProductService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)

	products, err := productService.GetAllProducts()
	if err != nil {
		fmt.Printf("Error finding all products: %v\n", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	for _, product := range products {
		table.Append([]string{
			fmt.Sprintf("%d", product.ProductID),
			product.Name,
			fmt.Sprintf("%.2f", product.Price),
			product.Description.String,
			product.Images.String,
			product.Type,
		})
	}
	table.Render()
	fmt.Println()
}

func findAllProductsMenu(db *sqlx.DB, productService *services.ProductService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find All Products")
	fmt.Println("=====================================")

	allProducts(productService)

	messages.PrintMessage(message)

	var input string
	fmt.Print("0. Back: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		findAllProductsMenu(db, productService, message)
	}

	if input == "0" {
		ManageProductMenu(db, "")
	} else {
		message = "Invalid input"
		messages.PrintMessage(message)
		findAllProductsMenu(db, productService, message)
	}
}

func findProductByNameMenu(db *sqlx.DB, productService *services.ProductService, message string) {
	terminal.Clear()

	fmt.Println("=====================================")
	fmt.Println("Find Product By Name")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Enter product name: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		findProductByNameMenu(db, productService, message)
	}

	products, err := productService.GetProductByName(input)
	if err != nil {
		fmt.Printf("Error finding product by name: %v\n", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)

	for _, product := range products {
		table.Append([]string{
			fmt.Sprintf("%d", product.ProductID),
			product.Name,
			fmt.Sprintf("%.2f", product.Price),
			product.Description.String,
			product.Images.String,
			product.Type,
		})
	}
	table.Render()
	fmt.Println()

	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}
	ManageProductMenu(db, "")
}
