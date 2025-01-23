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

func (service *PaymentMethodService) FindAll() ([]entity.PaymentMethod, error) {
	paymentMethods, err := service.paymentMethodRepository.FindAll()
	if err != nil {
		return paymentMethods, err
	}
	return paymentMethods, nil
}

func (service *PaymentMethodService) FindByID(paymentMethodID int) (entity.PaymentMethod, error) {
	paymentMethod, err := service.paymentMethodRepository.FindByID(paymentMethodID)
	if err != nil {
		return paymentMethod, err
	}
	return paymentMethod, nil
}

func (service *PaymentMethodService) Add(paymentMethod entity.PaymentMethod) error {
	err := service.paymentMethodRepository.Add(paymentMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *PaymentMethodService) Update(paymentMethod entity.PaymentMethod) error {
	err := service.paymentMethodRepository.Update(paymentMethod)
	if err != nil {
		return err
	}
	return nil
}

func (service *PaymentMethodService) Delete(paymentMethodID int) error {
	err := service.paymentMethodRepository.Delete(paymentMethodID)
	if err != nil {
		return err
	}
	return nil
}

func (service *PaymentMethodService) ResetIncrement() error {
	err := service.paymentMethodRepository.ResetIncrement()
	if err != nil {
		return err
	}
	return nil
}
