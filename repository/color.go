package repository

import "clothing-pair-project/entity"

type ColorsRepository interface {
	FindAll() ([]entity.Color, error)
	FindByID(colorID int) (entity.Color, error)
	Add(color entity.Color) error
	Update(color entity.Color) error
	Delete(colorID int) error
	ResetIncrement() error
}
