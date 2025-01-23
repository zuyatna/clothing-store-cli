package repository

import "clothing-pair-project/entity"

type CollectionRepository interface {
	FindAll() ([]entity.Collection, error)
	FindByID(collectionID int) (entity.Collection, error)
	Add(collection entity.Collection) error
	Update(collection entity.Collection) error
	Delete(collectionID int) error
	ResetIncrement() error
}
