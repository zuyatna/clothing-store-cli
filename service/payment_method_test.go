package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentMethodRepository = &repository.PaymentMethodRepositoryMock{Mock: mock.Mock{}}
var paymentMethodService = NewPaymentMethodService(paymentMethodRepository)

func TestFindAllPaymentMethods(t *testing.T) {
	paymentMethods := []entity.PaymentMethod{
		{
			PaymentMethodID: 1,
			Name:            "Credit Card",
		},
		{
			PaymentMethodID: 2,
			Name:            "Bank Transfer",
		},
	}

	paymentMethodRepository.On("FindAll").Return(paymentMethods, nil)

	result, err := paymentMethodRepository.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, paymentMethods, result)
	paymentMethodRepository.AssertExpectations(t)
}

func TestFindPaymentMethodByID(t *testing.T) {
	paymentMethod := entity.PaymentMethod{
		PaymentMethodID: 1,
		Name:            "Credit Card",
	}

	paymentMethodRepository.On("FindByID", 1).Return(paymentMethod, nil)

	result, err := paymentMethodRepository.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, paymentMethod, result)
	paymentMethodRepository.AssertExpectations(t)
}

func TestAddPaymentMethod(t *testing.T) {
	paymentMethod := entity.PaymentMethod{
		PaymentMethodID: 1,
		Name:            "Credit Card",
	}

	paymentMethodRepository.On("Add", paymentMethod).Return(nil)

	err := paymentMethodService.Add(paymentMethod)
	assert.NoError(t, err)
	paymentMethodRepository.AssertExpectations(t)
}

func TestUpdatePaymentMethod(t *testing.T) {
	paymentMethod := entity.PaymentMethod{
		PaymentMethodID: 1,
		Name:            "Credit Card",
	}

	paymentMethodRepository.On("Update", paymentMethod).Return(nil)

	err := paymentMethodService.Update(paymentMethod)
	assert.NoError(t, err)
	paymentMethodRepository.AssertExpectations(t)
}

func TestDeletePaymentMethod(t *testing.T) {
	paymentMethodRepository.On("Delete", 1).Return(nil)

	err := paymentMethodService.Delete(1)
	assert.NoError(t, err)
	paymentMethodRepository.AssertExpectations(t)
}

func TestResetIncrementPaymentMethod(t *testing.T) {
	paymentMethodRepository.On("ResetIncrement").Return(nil)

	err := paymentMethodService.ResetIncrement()
	assert.NoError(t, err)
	paymentMethodRepository.AssertExpectations(t)
}
