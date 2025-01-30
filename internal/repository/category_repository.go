package repository

import "clothing-pair-project/internal/models"

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindByID(id int) (models.Category, error)
	Add(category models.Category) error
	Update(category models.Category) error
	Delete(id int) error
}
