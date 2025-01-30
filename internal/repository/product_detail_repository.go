package repository

import "clothing-pair-project/internal/models"

type ProductDetailRepository interface {
	FindAll() ([]models.ProductDetail, error)
	FindByID(id int) (models.ProductDetail, error)
	Add(productDetail models.ProductDetail) error
	Update(productDetail models.ProductDetail) error
	Delete(id int) error
}
