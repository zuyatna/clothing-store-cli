package repository

import "clothing-pair-project/internal/models"

type ColorRepository interface {
	FindAll() ([]models.Color, error)
	FindByID(id int) (models.Color, error)
	Add(color models.Color) error
	Update(color models.Color) error
	Delete(id int) error
}
