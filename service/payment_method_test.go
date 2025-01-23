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

func TestGetAll(t *testing.T) {
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

	paymentMethodRepository.On("GetAll").Return(paymentMethods, nil)

	result, err := paymentMethodService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, paymentMethods, result)
	paymentMethodRepository.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	paymentMethod := entity.PaymentMethod{
		PaymentMethodID: 1,
		Name:            "Credit Card",
	}

	paymentMethodRepository.On("GetByID", 1).Return(paymentMethod, nil)

	result, err := paymentMethodService.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, paymentMethod, result)
	paymentMethodRepository.AssertExpectations(t)
}
