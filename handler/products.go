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
	query := `INSERT INTO products (category_id, color_id, size_id, price, stock, description, image, name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := h.db.Exec(query, products.Category_Id, products.Color_Id, products.Size_Id, products.Price, products.Stock, products.Description, products.Image, products.Name)
	return err
}

func (h *ProductMethodHandler) Delete(products int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := h.db.Exec(query, products)
	return err
}

func (h *ProductMethodHandler) Update(products entity.Products) error {
	query := `UPDATE products SET category_id = $1, color_id=$2, size_id=$3, price=$4, stock=$5, description=$6, image = $7, name = $8 WHERE product_id = $9`
	_, err := h.db.Exec(query, products.Category_Id, products.Color_Id, products.Size_Id, products.Price, products.Stock, products.Description, products.Image, products.Name, products.Product_Id)
	return err
}

func (h *ProductMethodHandler) Find(productID *int) ([]entity.ShowDataProducts, error) {
	var products []entity.ShowDataProducts
	var query string
	var err error

	if productID == nil {
		query = `select p.product_id, ca.name as category, co.name as color, s.name as size, p.name, p.price, p.stock, p.description, p.image, p.created_at 
		from products p, categories ca, colors co, sizes s where p.category_id=ca.category_id and p.color_id=co.color_id and p.size_id=s.size_id`
		err = h.db.Select(&products, query)
	} else {
		query = `select p.product_id, ca.name as category, co.name as color, s.name as size, p.name, p.price, p.stock, p.description, p.image, p.created_at 
		from products p, categories ca, colors co, sizes s where p.category_id=ca.category_id and p.color_id=co.color_id and p.size_id=s.size_id and product_id = $1`
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

func ShowDataProduct(namatable string, products []entity.ShowDataProducts) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(strings.Repeat(" ", 15) + namatable + strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("=", 40))
	_, _ = w.Write([]byte("ID\tCategory\tColor\tSize\tName\tPrice\tStock\tDesc\tImage\tCreatedAt\n"))
	_, _ = w.Write([]byte("--\t--\t--\t--\t--\t--\t--\t--\t--\t--\n"))

	for _, size := range products {
		_, _ = w.Write([]byte(
			fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%f\t%d\t%s\t%s\t%s\n", size.Product_Id, size.Category_Id, size.Color_Id, size.Size_Id, size.Name, size.Price, size.Stock, size.Description, size.Image, size.Created_At),
		))
	}

	_ = w.Flush()
	fmt.Println(strings.Repeat("=", 40))
}

func (h *ProductMethodHandler) AddProduct(product entity.Products) error {
	nextID, err := h.GetNextProductID()
	if err != nil {
		return err
	}

	query := `INSERT INTO products (product_id, category_id, color_id, size_id, name, price, stock, description, image, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP)`
	_, err = h.db.Exec(query, nextID, product.Category_Id, product.Color_Id, product.Size_Id, product.Name, product.Price, product.Stock, product.Description, product.Image)
	return err
}

func (h *ProductMethodHandler) GetNextProductID() (int, error) {
	var nextID int
	query := `SELECT setval('products_product_id_seq', (SELECT MAX(product_id)+1 FROM products));`
	err := h.db.Get(&nextID, query)
	return nextID, err
}

func (h *ProductMethodHandler) FindProductByName(name string) ([]entity.ShowDataProducts, error) {
	var products []entity.ShowDataProducts
	query := `select p.product_id, ca.name as category, co.name as color, s.name as size, p.name, p.price, p.stock, p.description, p.image, p.created_at 
	from products p, categories ca, colors co, sizes s where p.category_id=ca.category_id and p.color_id=co.color_id and p.size_id=s.size_id and p.name = $1`
	err := h.db.Select(&products, query, name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (h *ProductMethodHandler) UpdateProduct(product entity.Products) error {
	query := `UPDATE products SET category_id = $1, color_id = $2, size_id = $3, name = $4, price = $5, stock = $6, description = $7, image = $8 WHERE product_id = $9`
	_, err := h.db.Exec(query, product.Category_Id, product.Color_Id, product.Size_Id, product.Name, product.Price, product.Stock, product.Description, product.Image, product.Product_Id)
	return err
}

func (h *ProductMethodHandler) DeleteProduct(productID int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := h.db.Exec(query, productID)
	return err
}
