package menu

import (
	"bufio"
	"clothing-pair-project/internal/database/sqlrepo"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/internal/utils/messages"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	productRepository := sqlrepo.NewProductRepository(db)
	productService := services.NewProductService(productRepository)

	switch input {
	case "1":
		findAllProductsMenu(db, productService, "", 5, 0)
	case "2":
		findProductByNameMenu(db, productService, "")
	case "3":
		findProductByCategoryIDMenu(db, productService, "", 5, 0, false, 0)
	case "4":
		// Add Product
	case "5":
		editProductMenu(db, productService, "", 5, 0)
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

func showProducts(productService *services.ProductService, limit, offset int) (bool, bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Description", "Image", "Type"})
	table.SetRowLine(true)

	products, err := productService.GetAllProducts(limit+1, offset)
	if err != nil {
		fmt.Printf("Error finding all products: %v\n", err)
		return false, false
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return false, false
	}

	displayProducts := products
	if len(products) > limit {
		displayProducts = products[:limit]
	}

	for _, product := range displayProducts {
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

	return hasNext, hasPrev
}

func findAllProductsMenu(db *sqlx.DB, productService *services.ProductService, message string, limit, offset int) {
	fmt.Println("=====================================")
	fmt.Println("Find All Products")
	fmt.Println("=====================================")

	hasNext, hasPrev := showProducts(productService, limit, offset)

	messages.PrintMessage(message)

	var input string
	if hasPrev {
		fmt.Println("Type A to Previous")
	}
	if hasNext {
		fmt.Println("Type D to Next")
	}
	fmt.Println("Type 0 to Back")
	fmt.Println("Type Other Number to Choose Product")
	fmt.Print("Choose option: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}

	switch input {
	case "D", "d":
		if hasNext {
			offset += limit
		}
		findAllProductsMenu(db, productService, "", limit, offset)
	case "A", "a":
		if hasPrev {
			offset -= limit
			if offset < 0 {
				offset = 0
			}
		}
		findAllProductsMenu(db, productService, "", limit, offset)
	case "0":
		ManageProductMenu(db, "")
	default:
		findProductDetailByProductID(db, productService, limit, offset, input)
	}
}

func findProductDetailByProductID(db *sqlx.DB, productService *services.ProductService, lengthItem int, starItem int, input string) {
	productID, err := strconv.Atoi(input)
	if err != nil {
		msg := "Invalid product ID"
		messages.PrintMessage(msg)
		findAllProductsMenu(db, productService, msg, lengthItem, starItem)
		return
	}
	product, err := productService.GetProductByID(productID)
	if err != nil {
		msg := "Error finding product by ID"
		messages.PrintMessage(msg)
		findAllProductsMenu(db, productService, msg, lengthItem, starItem)
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

	productDetailRequestRepository := sqlrepo.NewProductDetailRequestRepository(db)
	productDetailRequestService := services.NewProductDetailRequestService(productDetailRequestRepository)

	productDetails, err := productDetailRequestService.GetProductDetailByProductID(productID)
	if err != nil {
		msg := "Error finding product detail by ID"
		messages.PrintMessage(msg)
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

	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		msg := "Error reading input"
		messages.PrintMessage(msg)
		ManageProductMenu(db, msg)
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

func findProductByCategoryIDMenu(db *sqlx.DB, productService *services.ProductService, message string, limit int, offset int, isActive bool, categoryID int) {
	fmt.Println("=====================================")
	fmt.Println("Find Product By Category")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	categoryRepository := sqlrepo.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	allCategories(categoryService)

	if !isActive {
		var input string
		fmt.Print("Enter category ID: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			findProductByCategoryIDMenu(db, productService, message, limit, offset, false, 0)
		}

		categoryID, err = strconv.Atoi(input)
		if err != nil {
			message = "Invalid category ID"
			messages.PrintMessage(message)
			findProductByCategoryIDMenu(db, productService, message, limit, offset, false, 0)
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

	var option string
	if hasPrev {
		fmt.Println("Type A to Previous")
	}
	if hasNext {
		fmt.Println("Type D to Next")
	}
	fmt.Println("Type 0 to Back")
	fmt.Println("Type Other Number to Choose Product")
	fmt.Print("Choose option: ")
	_, err = fmt.Scanln(&option)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		ManageProductMenu(db, message)
	}

	switch option {
	case "D", "d":
		if hasNext {
			offset += limit
		}
		findProductByCategoryIDMenu(db, productService, "", limit, offset, true, categoryID)
	case "A", "a":
		if hasPrev {
			offset -= limit
			if offset < 0 {
				offset = 0
			}
		}
		findProductByCategoryIDMenu(db, productService, "", limit, offset, true, categoryID)
	default:
		findProductDetailByProductID(db, productService, limit, offset, option)
	}
}

func editProductMenu(db *sqlx.DB, productService *services.ProductService, message string, limit, offset int) {
	fmt.Println("=====================================")
	fmt.Println("Edit Product")
	fmt.Println("=====================================")

	messages.PrintMessage(message)

	var input string
	fmt.Print("Enter product ID: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	productID, err := strconv.Atoi(input)
	if err != nil {
		message = "Invalid product ID"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	product, err := productService.GetProductByID(productID)
	if err != nil {
		message = "Error finding product by ID"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
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
		message = "No input entered"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	if input == "y" {

		categoryRepository := sqlrepo.NewCategoryRepository(db)
		categoryService := services.NewCategoryService(categoryRepository)

		categories, err := categoryService.GetCategoryByID(product.CategoryID)
		if err != nil {
			message = "Error finding category by ID"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Category: ", categories.Name)

		fmt.Print("Do you want to change the category? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			allCategories(categoryService)
			fmt.Print("Enter new category ID: ")
			_, err = fmt.Scanln(&input)
			if err != nil {
				message = "No input entered"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			newCategoryID, err := strconv.Atoi(input)
			if err != nil {
				message = "Invalid category ID"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.CategoryID = newCategoryID
		} else if input != "n" {
			message = "Invalid input"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Product Name: ", product.Name)

		fmt.Print("Do you want to change the product name? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			fmt.Print("Enter new product name: ")
			updateProductName, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductName = strings.TrimSpace(updateProductName)
			if err != nil {
				message = "Error reading input"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			if updateProductName == "" {
				message = "Product name cannot be empty"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.Name = updateProductName
		} else if input != "n" {
			message = "Invalid input"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Product Price: ", product.Price)

		fmt.Print("Do you want to change the product price? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			fmt.Print("Enter new product price: ")
			updateProductPrice, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductPrice = strings.TrimSpace(updateProductPrice)
			if err != nil {
				message = "Error reading input"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			if updateProductPrice == "" {
				message = "Product price cannot be empty"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			updateProductPriceFloat, err := strconv.ParseFloat(updateProductPrice, 64)
			if err != nil {
				message = "Invalid price"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			if updateProductPriceFloat < 0 {
				message = "Product price cannot be negative"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.Price = updateProductPriceFloat
		} else if input != "n" {
			message = "Invalid input"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Product Description: ", product.Description.String)

		fmt.Print("Do you want to change the product description? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			fmt.Print("Enter new product description: ")
			updateProductDescription, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductDescription = strings.TrimSpace(updateProductDescription)
			if err != nil {
				message = "Error reading input"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.Description = sql.NullString{
				String: updateProductDescription,
				Valid:  true,
			}
		} else if input != "n" {
			message = "Invalid input"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Product Image: ", product.Images.String)

		fmt.Print("Do you want to change the product image? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			fmt.Print("Enter new product image: ")
			updateProductImage, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductImage = strings.TrimSpace(updateProductImage)
			if err != nil {
				message = "Error reading input"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.Images = sql.NullString{
				String: updateProductImage,
				Valid:  true,
			}
		} else if input != "n" {
			message = "Invalid input"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		fmt.Println("Current Product Type: ", product.Type)

		fmt.Print("Do you want to change the product type? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			message = "No input entered"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		if input == "y" {
			// TODO: Implement change product type

			fmt.Print("Enter new product type: ")
			updateProductType, err := bufio.NewReader(os.Stdin).ReadString('\n')
			updateProductType = strings.TrimSpace(updateProductType)
			if err != nil {
				message = "Error reading input"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			if updateProductType == "" {
				message = "Product type cannot be empty"
				messages.PrintMessage(message)
				editProductMenu(db, productService, message, limit, offset)
			}

			product.Type = updateProductType
		}
	} else if input != "n" {
		message = "Invalid input"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
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
		message = "No input entered"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	if input == "y" {
		err = productService.UpdateProduct(product)
		if err != nil {
			message = "Error updating product"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
		}

		message = "Product updated successfully"
		messages.PrintMessage(message)
	} else if input != "n" {
		message = "Invalid input"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	fmt.Print("Do you want to edit the product detail? (y/n): ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		message = "No input entered"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

	if input == "y" {
		// TODO: Implement edit product detail

		productDetailRequestRepository := sqlrepo.NewProductDetailRequestRepository(db)
		productDetailRequestService := services.NewProductDetailRequestService(productDetailRequestRepository)

		productDetails, err := productDetailRequestService.GetProductDetailByProductID(productID)
		if err != nil {
			message = "Error finding product detail by ID"
			messages.PrintMessage(message)
			editProductMenu(db, productService, message, limit, offset)
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
		message = "Invalid input"
		messages.PrintMessage(message)
		editProductMenu(db, productService, message, limit, offset)
	}

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
