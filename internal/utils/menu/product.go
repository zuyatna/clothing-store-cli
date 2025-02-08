package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/handler"
	"clothing-pair-project/internal/utils/key_input"
	"clothing-pair-project/internal/utils/messages"
	"clothing-pair-project/internal/utils/tables"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
)

func ManageProductMenu(db *sqlx.DB, msg string) {
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

	messages.PrintMessage(msg)

	productRepository := sqlrepo.NewProductQuery(db)
	productService := services.NewProductService(productRepository)

	var input string
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		ManageUserMenu(db, msg)
		return
	}

	switch input {
	case "1":
		findAllProductsMenu(db, productService, "", 5, 0)
		return
	case "2":
		findProductByNameMenu(db, productService, "")
		return
	case "3":
		findProductByCategoryIDMenu(db, productService, "", 5, 0, false, 0)
		return
	case "4":
		// Add Product
		return
	case "5":
		editProductMenu(db, productService, "", 5, 0)
		return
	case "6":
		// Delete Product
		return
	case "0":
		AdminMenu(db, "")
		return
	default:
		msg = "Invalid key_input"
		ManageProductMenu(db, msg)
		return
	}
}

func showProducts(productService *services.ProductService, limit, offset int) (bool, bool) {
	table := tablewriter.NewWriter(os.Stdout)
	displayed := tables.ProductsTablePresenter(table)
	productHandler := handler.NewProductHandler(displayed, productService)

	hasNext, hasPrev, err := productHandler.ShowAllProducts(limit, offset)
	if err != nil {
		fmt.Printf("Error showing products: %v\n", err)
		return false, false
	}

	return hasNext, hasPrev
}

func showProductByID(productService *services.ProductService, productID int) {
	table := tablewriter.NewWriter(os.Stdout)
	displayed := tables.ProductsTablePresenter(table)
	productHandler := handler.NewProductHandler(displayed, productService)

	err := productHandler.ShowProductByID(productID)
	if err != nil {
		fmt.Printf("Error showing product by ID: %v\n", err)
		return
	}
}

func findAllProductsMenu(db *sqlx.DB, productService *services.ProductService, msg string, limit, offset int) {
	fmt.Println("=====================================")
	fmt.Println("Find All Products")
	fmt.Println("=====================================")

	hasNext, hasPrev := showProducts(productService, limit, offset)

	messages.PrintMessage(msg)

	if hasPrev {
		fmt.Println("Type A to Previous")
	}
	if hasNext {
		fmt.Println("Type D to Next")
	}
	fmt.Println("Type 0 to Back")
	fmt.Println("Type Other Number to Choose Product")
	fmt.Print("Choose option: ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No input entered"
		ManageProductMenu(db, msg)
		return
	}

	switch input {
	case "D", "d":
		if hasNext {
			offset += limit
		}
		findAllProductsMenu(db, productService, "", limit, offset)
		return
	case "A", "a":
		if hasPrev {
			offset -= limit
			if offset < 0 {
				offset = 0
			}
		}
		findAllProductsMenu(db, productService, "", limit, offset)
		return
	case "0":
		ManageProductMenu(db, "")
		return
	default:
		findProductDetailByProductID(db, productService, limit, offset, input)
		return
	}
}

func findProductDetailByProductID(db *sqlx.DB, productService *services.ProductService, lengthItem int, starItem int, input string) {
	productID, err := strconv.Atoi(input)
	if err != nil {
		msg := "Invalid product ID"
		findAllProductsMenu(db, productService, msg, lengthItem, starItem)
		return
	}

	fmt.Println()
	fmt.Println("The Product")
	fmt.Println()

	showProductByID(productService, productID)

	// TODO: Implement product detail first
	productDetailRequestRepository := sqlrepo.NewProductDetailRequestRepository(db)
	productDetailRequestService := services.NewProductDetailRequestService(productDetailRequestRepository)

	productDetails, err := productDetailRequestService.GetProductDetailByProductID(productID)
	if err != nil {
		msg := "Error finding product detail by ID"
		findAllProductsMenu(db, productService, msg, lengthItem, starItem)
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

	key_input.BackMenu()
	findAllProductsMenu(db, productService, "", lengthItem, starItem)
}

func findProductByNameMenu(db *sqlx.DB, productService *services.ProductService, msg string) {
	fmt.Println("=====================================")
	fmt.Println("Find Product By Name")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	var input string
	fmt.Print("Enter product name: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		findProductByNameMenu(db, productService, msg)
		return
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

	key_input.BackMenu()
	ManageProductMenu(db, "")
}

func findProductByCategoryIDMenu(db *sqlx.DB, productService *services.ProductService, msg string, limit int, offset int, isActive bool, categoryID int) {
	fmt.Println("=====================================")
	fmt.Println("Find Product By Category")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	categoryRepository := sqlrepo.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	allCategories(categoryService)

	if !isActive {
		var input string
		fmt.Print("Enter category ID: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			findProductByCategoryIDMenu(db, productService, msg, limit, offset, false, 0)
			return
		}

		categoryID, err = strconv.Atoi(input)
		if err != nil {
			msg = "Invalid category ID"
			messages.PrintMessage(msg)
			findProductByCategoryIDMenu(db, productService, msg, limit, offset, false, 0)
			return
		}
	}

	products, err := productService.GetProductByCategoryID(categoryID, limit+1, offset)
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

	displayProduct := products
	if len(products) > limit {
		displayProduct = products[:limit]
	}

	for _, product := range displayProduct {
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

	hasNext := len(products) > limit
	hasPrev := offset > 0

	if hasPrev {
		fmt.Println("Type A to Previous")
	}
	if hasNext {
		fmt.Println("Type D to Next")
	}
	fmt.Println("Type 0 to Back")
	fmt.Println("Type Other Number to Choose Product")
	fmt.Print("Choose option: ")

	var option string
	_, err = fmt.Scanln(&option)
	if err != nil {
		msg = "No key_input entered"
		ManageProductMenu(db, msg)
		return
	}

	switch option {
	case "D", "d":
		if hasNext {
			offset += limit
		}
		findProductByCategoryIDMenu(db, productService, "", limit, offset, true, categoryID)
		return
	case "A", "a":
		if hasPrev {
			offset -= limit
			if offset < 0 {
				offset = 0
			}
		}
		findProductByCategoryIDMenu(db, productService, "", limit, offset, true, categoryID)
		return
	default:
		findProductDetailByProductID(db, productService, limit, offset, option)
		return
	}
}

func editProductMenu(db *sqlx.DB, productService *services.ProductService, msg string, limit, offset int) {
	fmt.Println("=====================================")
	fmt.Println("Edit Product")
	fmt.Println("=====================================")

	messages.PrintMessage(msg)

	var input string
	fmt.Print("Enter product ID: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	productID, err := strconv.Atoi(input)
	if err != nil {
		msg = "Invalid product ID"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	product, err := productService.GetProductByID(productID)
	if err != nil {
		msg = "Error finding product by ID"
		editProductMenu(db, productService, msg, limit, offset)
		return
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

	fmt.Println()
	fmt.Print("Do you want to change the product? (y/n): ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	if input == "y" {
		categoryRepository := sqlrepo.NewCategoryRepository(db)
		categoryService := services.NewCategoryService(categoryRepository)

		categories, err := categoryService.GetCategoryByID(product.CategoryID)
		if err != nil {
			msg = "Error finding category by ID"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Category: ", categories.Name)
		fmt.Print("Do you want to change the category? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			allCategories(categoryService)
			fmt.Print("Enter new category ID: ")
			_, err = fmt.Scanln(&input)
			if err != nil {
				msg = "No key_input entered"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			newCategoryID, err := strconv.Atoi(input)
			if err != nil {
				msg = "Invalid category ID"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.CategoryID = newCategoryID
		} else if input != "n" {
			msg = "Invalid key_input"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Product Name: ", product.Name)
		fmt.Print("Do you want to change the product name? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			fmt.Print("Enter new product name: ")
			updateProductName, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductName = strings.TrimSpace(updateProductName)
			if err != nil {
				msg = "Error reading key_input"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			if updateProductName == "" {
				msg = "Product name cannot be empty"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.Name = updateProductName
		} else if input != "n" {
			msg = "Invalid key_input"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Product Price: ", product.Price)
		fmt.Print("Do you want to change the product price? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			fmt.Print("Enter new product price: ")
			updateProductPrice, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductPrice = strings.TrimSpace(updateProductPrice)
			if err != nil {
				msg = "Error reading key_input"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			if updateProductPrice == "" {
				msg = "Product price cannot be empty"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			updateProductPriceFloat, err := strconv.ParseFloat(updateProductPrice, 64)
			if err != nil {
				msg = "Invalid price"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			if updateProductPriceFloat < 0 {
				msg = "Product price cannot be negative"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.Price = updateProductPriceFloat
		} else if input != "n" {
			msg = "Invalid key_input"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Product Description: ", product.Description.String)
		fmt.Print("Do you want to change the product description? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			fmt.Print("Enter new product description: ")
			updateProductDescription, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductDescription = strings.TrimSpace(updateProductDescription)
			if err != nil {
				msg = "Error reading key_input"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.Description = sql.NullString{
				String: updateProductDescription,
				Valid:  true,
			}
		} else if input != "n" {
			msg = "Invalid key_input"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Product Image: ", product.Images.String)
		fmt.Print("Do you want to change the product image? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			fmt.Print("Enter new product image: ")
			updateProductImage, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductImage = strings.TrimSpace(updateProductImage)
			if err != nil {
				msg = "Error reading key_input"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.Images = sql.NullString{
				String: updateProductImage,
				Valid:  true,
			}
		} else if input != "n" {
			msg = "Invalid key_input"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		fmt.Println("Current Product Type: ", product.Type)
		fmt.Print("Do you want to change the product type? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			msg = "No key_input entered"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		if input == "y" {
			// TODO: Implement change product type

			fmt.Print("Enter new product type: ")
			updateProductType, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductType = strings.TrimSpace(updateProductType)
			if err != nil {
				msg = "Error reading key_input"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			if updateProductType == "" {
				msg = "Product type cannot be empty"
				editProductMenu(db, productService, msg, limit, offset)
				return
			}

			product.Type = updateProductType
		}
	} else if input != "n" {
		msg = "Invalid key_input"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	fmt.Println()
	fmt.Println("Updated Product")
	fmt.Println()

	table = tablewriter.NewWriter(os.Stdout)
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

	fmt.Println()
	fmt.Print("Do you want to save the changes? (y/n): ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	if input == "y" {
		err = productService.UpdateProduct(product)
		if err != nil {
			msg = "Error updating product"
			editProductMenu(db, productService, msg, limit, offset)
			return
		}

		msg = "Product updated successfully"
		messages.PrintMessage(msg)
	} else if input != "n" {
		msg = "Invalid key_input"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	fmt.Print("Do you want to edit the product detail? (y/n): ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		msg = "No key_input entered"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	if input == "y" {
		// TODO: Implement edit product detail

		productDetailRequestRepository := sqlrepo.NewProductDetailRequestRepository(db)
		productDetailRequestService := services.NewProductDetailRequestService(productDetailRequestRepository)

		productDetails, err := productDetailRequestService.GetProductDetailByProductID(productID)
		if err != nil {
			msg = "Error finding product detail by ID"
			editProductMenu(db, productService, msg, limit, offset)
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
	} else if input != "n" {
		msg = "Invalid key_input"
		editProductMenu(db, productService, msg, limit, offset)
		return
	}

	key_input.BackMenu()
	ManageProductMenu(db, "")
}
