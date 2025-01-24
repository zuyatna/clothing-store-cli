package repository

import "clothing-pair-project/entity"

type SizeRepository interface {
	Find(sizeID *int) ([]entity.Sizes, error)
	Add(size entity.Sizes) error
	Update(size entity.Sizes) error
	Delete(sizeID int) error
}
