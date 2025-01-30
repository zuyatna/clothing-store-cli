package repository

import "clothing-pair-project/internal/models"

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id int) (models.Product, error)
	FindByName(name string) ([]models.Product, error)
	FindByCategoryID(categoryID int) ([]models.Product, error)
	Add(product models.Product) error
	Update(product models.Product) error
	Delete(id int) error
}
