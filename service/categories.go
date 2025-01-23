package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type CategoryMethodService struct {
	sizeRepository repository.CategoryRepository
}

func NewCategoryMethodService(categoryMethodRepository repository.CategoryRepository) *CategoryMethodService {
	return &CategoryMethodService{categoryMethodRepository}
}

func (service *CategoryMethodService) Add(categoryMethod entity.Categories) error {
	err := service.sizeRepository.Add(categoryMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *CategoryMethodService) Update(categoryMethod entity.Categories) error {
	err := service.sizeRepository.Update(categoryMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *CategoryMethodService) Delete(categoryMethodID int) error {
	err := service.sizeRepository.Delete(categoryMethodID)
	if err != nil {
		return err
	}
	return nil
}

func (service *CategoryMethodService) Find(categoryMethodID *int) ([]entity.Categories, error) {
	size, err := service.sizeRepository.Find(categoryMethodID)
	if err != nil {
		return size, err
	}
	return size, nil
}
