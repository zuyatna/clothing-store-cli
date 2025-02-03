package sqlrepo

import (
	"clothing-pair-project/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) FindAll(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := "SELECT * FROM products ORDER BY product_id ASC LIMIT $1 OFFSET $2"
	err := repository.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *ProductRepository) FindByID(id int) (models.Product, error) {
	var product models.Product
	query := "SELECT * FROM products WHERE product_id = $1"
	err := repository.db.Get(&product, query, id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepository) FindByName(name string) ([]models.Product, error) {
	var products []models.Product
	query := "SELECT * FROM products WHERE name ILIKE $1"
	err := repository.db.Select(&products, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *ProductRepository) FindByCategoryID(categoryID int, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := "SELECT * FROM products WHERE category_id = $1 ORDER BY product_id ASC LIMIT $2 OFFSET $3"
	err := repository.db.Select(&products, query, categoryID, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *ProductRepository) Add(product models.Product) error {
	query := `INSERT INTO products (category_id, name, price, description, images, type, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`

	_, err := repository.db.Exec(query, product.CategoryID, product.Name, product.Price, product.Description, product.Images, product.Type)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepository) Update(product models.Product) error {
	query := `UPDATE products SET category_id=$1, name=$2, price=$3, description=$4, images=$5, type=$6 
			  WHERE product_id = $7`

	_, err := repository.db.Exec(query, product.CategoryID, product.Name, product.Price, product.Description, product.Images, product.Type, product.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE product_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
