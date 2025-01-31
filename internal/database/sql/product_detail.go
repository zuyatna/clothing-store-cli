package sql

import (
	"clothing-pair-project/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductDetailRepository struct {
	db *sqlx.DB
}

type ProductDetailRequestRepository struct {
	db *sqlx.DB
}

func NewProductDetailRepository(db *sqlx.DB) *ProductDetailRepository {
	return &ProductDetailRepository{db: db}
}

func NewProductDetailRequestRepository(db *sqlx.DB) *ProductDetailRequestRepository {
	return &ProductDetailRequestRepository{db: db}
}

func (repository *ProductDetailRepository) FindAll() ([]models.ProductDetail, error) {
	var productDetails []models.ProductDetail
	query := "SELECT * FROM product_details ORDER BY product_detail_id ASC"
	err := repository.db.Select(&productDetails, query)
	if err != nil {
		return nil, err
	}

	return productDetails, nil
}

func (repository *ProductDetailRepository) FindByID(id int) (models.ProductDetail, error) {
	var productDetail models.ProductDetail
	query := "SELECT * FROM product_details WHERE product_detail_id = $1"
	err := repository.db.Get(&productDetail, query, id)
	if err != nil {
		return models.ProductDetail{}, err
	}

	return productDetail, nil
}

func (repository *ProductDetailRepository) FindByProductID(productID int) ([]models.ProductDetail, error) {
	var productDetails []models.ProductDetail
	query := "SELECT * FROM product_details WHERE product_id = $1 ORDER BY product_detail_id ASC"
	err := repository.db.Select(&productDetails, query, productID)
	if err != nil {
		return nil, err
	}

	return productDetails, nil
}

func (repository *ProductDetailRepository) Add(productDetail models.ProductDetail) error {
	nextID, err := repository.GetNextID()
	if err != nil {
		return err
	}

	query := `INSERT INTO product_details (product_detail_id, product_id, color_id, size_id, stock, created_at) 
			  VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)`

	_, err = repository.db.Exec(query, nextID, productDetail.ProductID, productDetail.ColorID, productDetail.SizeID, productDetail.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductDetailRepository) Update(productDetail models.ProductDetail) error {
	query := "UPDATE product_details SET product_id = $1, color_id = $2, size_id = $3, stock = $4 WHERE product_detail_id = $5"
	_, err := repository.db.Exec(query, productDetail.ProductID, productDetail.ColorID, productDetail.SizeID, productDetail.Stock, productDetail.ProductDetailID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductDetailRepository) Delete(id int) error {
	query := "DELETE FROM product_details WHERE product_detail_id = $1"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductDetailRepository) GetNextID() (int, error) {
	var id int
	query := "SELECT MAX(product_detail_id) FROM product_details"
	err := repository.db.Get(&id, query)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (repository *ProductDetailRequestRepository) FindProductDetailByID(productID int) ([]models.ProductDetailRequest, error) {
	var productDetailRequests []models.ProductDetailRequest
	query := `SELECT c.name AS color, s.name AS size, pd.stock
          FROM product_details pd
          JOIN colors c ON pd.color_id = c.color_id
          JOIN sizes s ON pd.size_id = s.size_id
          WHERE pd.product_id = $1
          ORDER BY pd.product_detail_id ASC`
	err := repository.db.Select(&productDetailRequests, query, productID)
	if err != nil {
		return nil, err
	}

	return productDetailRequests, nil
}
