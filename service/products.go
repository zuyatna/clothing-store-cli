package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type ProductMethodService struct {
	productRepository repository.ProductRepository
}

func NewProductMethodService(productMethodRepository repository.ProductRepository) *ProductMethodService {
	return &ProductMethodService{productMethodRepository}
}

func (service *ProductMethodService) Add(productMethod entity.Products) error {
	err := service.productRepository.Add(productMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductMethodService) Update(productMethod entity.Products) error {
	err := service.productRepository.Update(productMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductMethodService) Delete(productMethodID int) error {
	err := service.productRepository.Delete(productMethodID)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductMethodService) Find(productMethodID *int) ([]entity.ShowDataProducts, error) {
	product, err := service.productRepository.Find(productMethodID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (service *ProductMethodService) AddProduct(productMethod entity.Products) error {
	err := service.productRepository.AddProduct(productMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductMethodService) FindProductByName(productName string) ([]entity.ShowDataProducts, error) {
	product, err := service.productRepository.FindProductByName(productName)
	if err != nil {
		return product, err
	}
	return product, nil

}

func (service *ProductMethodService) UpdateProduct(productMethod entity.Products) error {
	err := service.productRepository.UpdateProduct(productMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductMethodService) DeleteProduct(productMethodID int) error {
	err := service.productRepository.DeleteProduct(productMethodID)
	if err != nil {
		return err
	}
	return nil
}
