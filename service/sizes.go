package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type SizeMethodService struct {
	sizeRepository repository.SizeRepository
}

func NewSizeMethodService(sizeMethodRepository repository.SizeRepository) *SizeMethodService {
	return &SizeMethodService{sizeMethodRepository}
}

func (service *SizeMethodService) Add(sizeMethod entity.Sizes) error {
	err := service.sizeRepository.Add(sizeMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *SizeMethodService) Update(sizeMethod entity.Sizes) error {
	err := service.sizeRepository.Update(sizeMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *SizeMethodService) Delete(sizeMethodID int) error {
	err := service.sizeRepository.Delete(sizeMethodID)
	if err != nil {
		return err
	}
	return nil
}

func (service *SizeMethodService) Find(sizeMethodID *int) ([]entity.Sizes, error) {
	size, err := service.sizeRepository.Find(sizeMethodID)
	if err != nil {
		return size, err
	}
	return size, nil
}
