package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type CollectionService struct {
	collectionRepository repository.CollectionRepository
}

func NewCollectionService(collectionRepository repository.CollectionRepository) *CollectionService {
	return &CollectionService{collectionRepository}
}

func (service *CollectionService) FindAll() ([]entity.Collection, error) {
	collections, err := service.collectionRepository.FindAll()
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (service *CollectionService) FindByID(collectionID int) (entity.Collection, error) {
	collection, err := service.collectionRepository.FindByID(collectionID)
	if err != nil {
		return collection, err
	}
	return collection, nil
}

func (service *CollectionService) Add(collection entity.Collection) error {
	err := service.collectionRepository.Add(collection)
	if err != nil {
		return err
	}
	return nil
}

func (service *CollectionService) Update(collection entity.Collection) error {
	err := service.collectionRepository.Update(collection)
	if err != nil {
		return err
	}
	return nil
}

func (service *CollectionService) Delete(collectionID int) error {
	err := service.collectionRepository.Delete(collectionID)
	if err != nil {
		return err
	}
	return nil
}

func (service *CollectionService) ResetIncrement() error {
	err := service.collectionRepository.ResetIncrement()
	if err != nil {
		return err
	}
	return nil
}
