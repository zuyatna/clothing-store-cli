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

func ManageCheckoutMenu(db *sqlx.DB) {
	checkoutHandler := handler.NewProductsHandler(db)
	checkoutService := service.NewProductMethodService(checkoutHandler)
	chooseProduct(db, checkoutService)
}

func allProductReady(productService *service.ProductMethodService) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Category", "Color", "Size", "Name", "Description", "Price", "Stock"})

	products, err := productService.Find(nil)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No product found.")
		return
	}

	for _, product := range products {
		table.Append([]string{
			strconv.Itoa(product.Product_Id),
			product.Category_Id,
			product.Color_Id,
			product.Size_Id,
			product.Name,
			product.Description,
			strconv.FormatFloat(float64(product.Price), 'f', 2, 32),
			strconv.Itoa(product.Stock),
		})
	}
	table.Render()

	fmt.Println()
}

func chooseProduct(db *sqlx.DB, productService *service.ProductMethodService) {
	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("Choose Product")
	fmt.Println("=====================================")

	allProductReady(productService)

	var input string
	fmt.Print("0. Back: ")
	fmt.Scanln(&input)

	if input == "0" {
		return
	}

	productID, err := strconv.Atoi(input) // Ubah input string menjadi integer
	if err != nil {
		fmt.Println("Invalid input, please enter a valid product ID.")
		return
	}

	id := &productID
	products, err := productService.Find(id)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if len(products) == 0 {
		fmt.Println("Invalid input")
	}

	ManagePaymentMethodMenu(db)
}
