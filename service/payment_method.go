package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type PaymentMethodService struct {
	paymentMethodRepository repository.PaymentMethodRepository
}

func NewPaymentMethodService(paymentMethodRepository repository.PaymentMethodRepository) *PaymentMethodService {
	return &PaymentMethodService{paymentMethodRepository}
}

func (service *PaymentMethodService) GetAll() ([]entity.PaymentMethod, error) {
	paymentMethods, err := service.paymentMethodRepository.GetAll()
	if err != nil {
		return paymentMethods, err
	}
	return paymentMethods, nil
}

func (service *PaymentMethodService) GetByID(paymentMethodID int) (entity.PaymentMethod, error) {
	paymentMethod, err := service.paymentMethodRepository.GetByID(paymentMethodID)
	if err != nil {
		return paymentMethod, err
	}
	return paymentMethod, nil
}
