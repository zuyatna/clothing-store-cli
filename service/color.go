package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type ColorService struct {
	colorRepository repository.ColorsRepository
}

func NewColorService(colorRepository repository.ColorsRepository) *ColorService {
	return &ColorService{colorRepository}
}

func (service *ColorService) FindAll() ([]entity.Color, error) {
	colors, err := service.colorRepository.FindAll()
	if err != nil {
		return colors, err
	}
	return colors, nil
}

func (service *ColorService) FindByID(colorID int) (entity.Color, error) {
	color, err := service.colorRepository.FindByID(colorID)
	if err != nil {
		return color, err
	}
	return color, nil
}

func (service *ColorService) Add(color entity.Color) error {
	err := service.colorRepository.Add(color)
	if err != nil {
		return err
	}
	return nil
}

func (service *ColorService) Update(color entity.Color) error {
	err := service.colorRepository.Update(color)
	if err != nil {
		return err
	}
	return nil
}

func (service *ColorService) Delete(colorID int) error {
	err := service.colorRepository.Delete(colorID)
	if err != nil {
		return err
	}
	return nil
}

func (service *ColorService) ResetIncrement() error {
	err := service.colorRepository.ResetIncrement()
	if err != nil {
		return err
	}
	return nil
}
