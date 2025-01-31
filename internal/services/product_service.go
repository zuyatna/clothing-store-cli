package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (service *ProductService) GetAllProducts(limit, offset int) ([]models.Product, error) {
	return service.productRepository.FindAll(limit, offset)
}

func (service *ProductService) GetProductByID(id int) (models.Product, error) {
	return service.productRepository.FindByID(id)
}

func (service *ProductService) GetProductByName(name string) ([]models.Product, error) {
	return service.productRepository.FindByName(name)
}

func (service *ProductService) GetProductByCategoryID(categoryID int) ([]models.Product, error) {
	return service.productRepository.FindByCategoryID(categoryID)
}

func (service *ProductService) AddProduct(product models.Product) error {
	return service.productRepository.Add(product)
}

func (service *ProductService) UpdateProduct(product models.Product) error {
	return service.productRepository.Update(product)
}

func (service *ProductService) DeleteProduct(id int) error {
	return service.productRepository.Delete(id)
}
