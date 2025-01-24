package menu

import (
	"bufio"
	"clothing-pair-project/entity"
	"clothing-pair-project/handler"
	"clothing-pair-project/service"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageProductMenu(db *sqlx.DB) {
	for {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Manage Product Menu")
		fmt.Println("1. Add Product")
		fmt.Println("2. Find All Products")
		fmt.Println("3. Find Product By Name")
		fmt.Println("4. Update Product")
		fmt.Println("5. Delete Product")
		fmt.Println("0. Back")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose option: ")
		fmt.Scanln(&input)

		productHandler := handler.NewProductsHandler(db)
		productService := service.NewProductMethodService(productHandler)

		switch input {
		case 1:
			addProductMenu(productService, db)
		case 2:
			findAllProductsMenu(productService)
		case 3:
			// TODO: Find product by name menu
		case 4:
			// TODO: Update product menu
		case 5:
			// TODO: Delete product menu
		case 0:
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func addProductMenu(productService *service.ProductMethodService, db *sqlx.DB) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Add Product")
	fmt.Println("=====================================")

	var name, description, image string
	var categoryID, colorID, sizeID, stock int
	var price float32

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	colorHandler := handler.NewColorHandler(db)
	colorService := service.NewColorService(colorHandler)
	AllColor(colorService)

	fmt.Print("Color ID: ")
	fmt.Scanln(&colorID)

	sizeHandler := handler.NewSizesHandler(db)
	sizeService := service.NewSizeMethodService(sizeHandler)
	AllSizes(sizeService)

	fmt.Print("Size ID: ")
	fmt.Scanln(&sizeID)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Description: ")
	description, _ = reader.ReadString('\n')
	description = strings.TrimSpace(description)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	priceFloat, _ := strconv.ParseFloat(priceStr, 32)
	price = float32(priceFloat)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Stock: ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	stock, _ = strconv.Atoi(stockStr)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Image: ")
	image, _ = reader.ReadString('\n')
	image = strings.TrimSpace(image)

	categoryHandler := handler.NewCategoryHandler(db)
	categoryService := service.NewCategoryMethodService(categoryHandler)
	AllCategory(categoryService)

	fmt.Print("Category ID: ")
	fmt.Scanln(&categoryID)

	product := entity.Products{
		Category_Id: categoryID,
		Color_Id:    colorID,
		Size_Id:     sizeID,
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
		Image:       image,
	}

	err := productService.AddProduct(product)
	if err != nil {
		fmt.Printf("Failed to add product: %v\n", err)
		return
	}
	fmt.Println("Product added successfully!")
}

func findAllProductsMenu(productService *service.ProductMethodService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Find All Products")
	fmt.Println("=====================================")

	AllProducts(productService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)
	if input == "0" {
		return
	} else {
		fmt.Println("Invalid input")
	}
}

func AllProducts(productService *service.ProductMethodService) {
	products, err := productService.Find(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Color", "Size", "Name", "Price", "Stock", "Desc"})
	for _, product := range products {
		table.Append([]string{product.Category_Id, product.Color_Id, product.Size_Id, product.Name, strconv.FormatFloat(float64(product.Price), 'f', 2, 32), strconv.Itoa(product.Stock), product.Description})
	}
	table.Render()

	fmt.Println()
}
