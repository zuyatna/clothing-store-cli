package repository

import "clothing-pair-project/entity"

type ProductRepository interface {
	Find(productID *int) ([]entity.ShowDataProducts, error)
	Add(product entity.Products) error
	Update(product entity.Products) error
	Delete(productID int) error
}
