package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sql"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/messages"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageProductMenu(db *sqlx.DB, message string) {
	fmt.Println("=====================================")
	fmt.Println("Manage Product Menu")
	fmt.Println("1. Find All Products")
	fmt.Println("2. Find Product By Name")
	fmt.Println("3. Find Product By Category")
	fmt.Println("4. Add Product")
	fmt.Println("5. Edit Product")
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
		findAllProductsMenu(db, productService, "", 5, 0)
	case "2":
		findProductByNameMenu(db, productService, "")
	case "3":
		findProductByCategoryIDMenu(db, productService, "")
	case "4":
		// Add Product
	case "5":
		// edit Product
	case "6":
		// Delete Product
	case "0":
		AdminMenu(db, "")
	default:
		message = "Invalid input"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}
}

func showProducts(productService *services.ProductService, next, prev int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)

	products, err := productService.GetAllProducts(next, prev)
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

func findAllProductsMenu(db *sqlx.DB, productService *services.ProductService, message string, lengthItem, starItem int) {
	fmt.Println("=====================================")
	fmt.Println("Find All Products")
	fmt.Println("=====================================")

	showProducts(productService, lengthItem, starItem)

	messages.PrintMessage(message)

	products, err := productService.GetAllProducts(lengthItem, starItem)
	if err != nil {
		fmt.Printf("Error finding all products: %v\n", err)
		return
	}

	var input string
	if len(products) >= 6 {
		fmt.Println("Type A to Previous")
		fmt.Println("Type D to Next")
	}

	fmt.Println("Type 0 to Back")
	fmt.Println("Type Other Number to Choose Product")
	fmt.Print("Choose option: ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageUserMenu(db, message)
	}

	switch input {
	case "D", "d":
		starItem += lengthItem
		findAllProductsMenu(db, productService, "", lengthItem, starItem)
	case "A", "a":
		starItem -= lengthItem
		if starItem < 0 {
			starItem = 0
		}
		findAllProductsMenu(db, productService, "", lengthItem, starItem)
	case "0":
		ManageProductMenu(db, "")
	default:
		findProductDetailByProductID(db, productService, message, lengthItem, starItem, input)
	}
}

func findProductDetailByProductID(db *sqlx.DB, productService *services.ProductService, message string, lengthItem int, starItem int, input string) {
	productID, err := strconv.Atoi(input)
	if err != nil {
		message = "Invalid product ID"
		messages.PrintMessage(message)
		findAllProductsMenu(db, productService, message, lengthItem, starItem)
	}

	product, err := productService.GetProductByID(productID)
	if err != nil {
		message = "Error finding product by ID"
		messages.PrintMessage(message)
		findAllProductsMenu(db, productService, message, lengthItem, starItem)
	}

	fmt.Println()
	fmt.Println("The Product")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)

	table.Append([]string{
		fmt.Sprintf("%d", product.ProductID),
		product.Name,
		fmt.Sprintf("%.2f", product.Price),
		product.Description.String,
		product.Images.String,
		product.Type,
	})
	table.Render()

	productDetailRequestRepository := sql.NewProductDetailRequestRepository(db)
	productDetailRequestService := services.NewProductDetailRequestService(productDetailRequestRepository)

	productDetails, err := productDetailRequestService.GetProductDetailByProductID(productID)
	if err != nil {
		message = "Error finding product detail by ID"
		messages.PrintMessage(message)
		findAllProductsMenu(db, productService, message, lengthItem, starItem)
		return
	}

	fmt.Println()
	fmt.Println("Detail of Products")
	fmt.Println()

	tableDetail := tablewriter.NewWriter(os.Stdout)
	tableDetail.SetHeader([]string{"No", "Color", "Size", "Stock"})
	tableDetail.SetRowLine(true)

	count := 0
	for _, detail := range productDetails {
		count++
		tableDetail.Append([]string{
			fmt.Sprintf("%d", count),
			detail.Color,
			detail.Size,
			fmt.Sprintf("%d", detail.Stock),
		})
	}
	tableDetail.Render()

	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		message = "Error reading input"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}

	findAllProductsMenu(db, productService, "", lengthItem, starItem)
}

func findProductByNameMenu(db *sqlx.DB, productService *services.ProductService, message string) {
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

func findProductByCategoryIDMenu(db *sqlx.DB, productService *services.ProductService, message string) {
	fmt.Println("=====================================")
	fmt.Println("Find Product By Category")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	categoryRepository := sql.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	allCategories(categoryService)

	var input string
	fmt.Print("Enter category ID: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		findProductByCategoryIDMenu(db, productService, message)
	}

	categoryID, err := strconv.Atoi(input)
	if err != nil {
		message = "Invalid category ID"
		messages.PrintMessage(message)
		findProductByCategoryIDMenu(db, productService, message)
		return
	}

	products, err := productService.GetProductByCategoryID(categoryID)
	if err != nil {
		fmt.Printf("Error finding product by category ID: %v\n", err)
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
