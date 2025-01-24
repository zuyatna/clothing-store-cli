package repository

import "clothing-pair-project/entity"

type CategoryRepository interface {
	Find(catID *int) ([]entity.Categories, error)
	Add(cat entity.Categories) error
	Update(cat entity.Categories) error
	Delete(catID int) error
}
