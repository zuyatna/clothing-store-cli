package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type SizeService struct {
	sizeRepository repository.SizeRepository
}

func NewSizeService(sizeRepository repository.SizeRepository) *SizeService {
	return &SizeService{sizeRepository: sizeRepository}
}

func (service *SizeService) GetAllSizes() ([]models.Size, error) {
	return service.sizeRepository.FindAll()
}

func (service *SizeService) GetSizeByID(id int) (models.Size, error) {
	return service.sizeRepository.FindByID(id)
}

func (service *SizeService) AddSize(size models.Size) error {
	return service.sizeRepository.Add(size)
}

func (service *SizeService) UpdateSize(size models.Size) error {
	return service.sizeRepository.Update(size)
}

func (service *SizeService) DeleteSize(id int) error {
	return service.sizeRepository.Delete(id)
}
