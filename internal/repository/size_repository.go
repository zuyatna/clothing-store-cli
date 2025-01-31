package repository

import "clothing-pair-project/internal/models"

type SizeRepository interface {
	FindAll() ([]models.Size, error)
	FindByID(id int) (models.Size, error)
	Add(size models.Size) error
	Update(size models.Size) error
	Delete(id int) error
}
