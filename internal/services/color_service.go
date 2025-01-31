package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type ColorService struct {
	colorRepository repository.ColorRepository
}

func NewColorService(colorRepository repository.ColorRepository) *ColorService {
	return &ColorService{colorRepository: colorRepository}
}

func (service *ColorService) GetAllColors() ([]models.Color, error) {
	return service.colorRepository.FindAll()
}

func (service *ColorService) GetColorByID(id int) (models.Color, error) {
	return service.colorRepository.FindByID(id)
}

func (service *ColorService) AddColor(color models.Color) error {
	return service.colorRepository.Add(color)
}

func (service *ColorService) UpdateColor(color models.Color) error {
	return service.colorRepository.Update(color)
}

func (service *ColorService) DeleteColor(id int) error {
	return service.colorRepository.Delete(id)
}
