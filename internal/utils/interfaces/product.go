package interfaces

import "clothing-pair-project/internal/models"

type ProductDisplay interface {
	DisplayProducts(products []models.Product)
	DisplayProduct(product models.Product)
}

type ProductFetcher interface {
	GetAllProducts(limit, offset int) ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
}
