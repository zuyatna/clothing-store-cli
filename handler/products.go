package handler

import (
	"clothing-pair-project/entity"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/jmoiron/sqlx"
)

type ProductMethodHandler struct {
	db *sqlx.DB
}

func NewProductsHandler(db *sqlx.DB) *ProductMethodHandler {
	return &ProductMethodHandler{db: db}
}

func (h *ProductMethodHandler) Add(products entity.Products) error {
	query := `INSERT INTO products (category_id, color_id, size_id, price, stock, description, image) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := h.db.Exec(query, products.Category_Id, products.Color_Id, products.Size_Id, products.Price, products.Stock, products.Description, products.Image)
	return err
}

func (h *ProductMethodHandler) Delete(products int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := h.db.Exec(query, products)
	return err
}

func (h *ProductMethodHandler) Update(products entity.Products) error {
	query := `UPDATE products SET category_id = $1, color_id=$2, size_id=$3, price=$4, stock=$5, description=$6, image = $7 WHERE product_id = $8`
	_, err := h.db.Exec(query, products.Category_Id, products.Color_Id, products.Size_Id, products.Price, products.Stock, products.Description, products.Image, products.Product_Id)
	return err
}

func (h *ProductMethodHandler) Find(productID *int) ([]entity.Products, error) {
	var products []entity.Products
	var query string
	var err error

	if productID == nil {
		query = `SELECT * FROM products`
		err = h.db.Select(&products, query)
	} else {
		query = `SELECT * FROM products WHERE product_id = $1`
		err = h.db.Select(&products, query, *productID)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	if err != nil {
		return nil, err
	}
	return products, nil
}

func ShowDataProduct(namatable string, products []entity.Products) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("ID\tCategory\tColor\tSize\tPrice\tStock\tDesc\tImage\tCreatedAt\n"))
	_, _ = w.Write([]byte("--\t--\t--\t--\t--\t--\t--\t--\t--\n"))

	for _, size := range products {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%d\t%d\t%d\t%d\t%f\t%d\t%s\t%s\t%s\n", size.Product_Id, size.Category_Id, size.Color_Id, size.Size_Id, size.Price, size.Stock, size.Description, size.Image, size.Created_At),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
}
