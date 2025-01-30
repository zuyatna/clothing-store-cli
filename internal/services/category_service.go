package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (service *CategoryService) GetAllCategories() ([]models.Category, error) {
	return service.categoryRepository.FindAll()
}

func (service *CategoryService) GetCategoryByID(id int) (models.Category, error) {
	return service.categoryRepository.FindByID(id)
}

func (service *CategoryService) AddCategory(category models.Category) error {
	return service.categoryRepository.Add(category)
}

func (service *CategoryService) UpdateCategory(category models.Category) error {
	return service.categoryRepository.Update(category)
}

func (service *CategoryService) DeleteCategory(id int) error {
	return service.categoryRepository.Delete(id)
}
