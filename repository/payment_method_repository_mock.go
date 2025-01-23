package repository

import (
	"clothing-pair-project/entity"

	"github.com/stretchr/testify/mock"
)

type PaymentMethodRepositoryMock struct {
	mock.Mock
}

func (m *PaymentMethodRepositoryMock) FindAll() ([]entity.PaymentMethod, error) {
	args := m.Called()
	return args.Get(0).([]entity.PaymentMethod), args.Error(1)
}

func (m *PaymentMethodRepositoryMock) FindByID(paymentMethodID int) (entity.PaymentMethod, error) {
	args := m.Called(paymentMethodID)
	return args.Get(0).(entity.PaymentMethod), args.Error(1)
}

func (m *PaymentMethodRepositoryMock) Add(paymentMethod entity.PaymentMethod) error {
	args := m.Called(paymentMethod)
	return args.Error(0)
}

func (m *PaymentMethodRepositoryMock) Update(paymentMethod entity.PaymentMethod) error {
	args := m.Called(paymentMethod)
	return args.Error(0)
}

func (m *PaymentMethodRepositoryMock) Delete(paymentMethodID int) error {
	args := m.Called(paymentMethodID)
	return args.Error(0)
}

func (m *PaymentMethodRepositoryMock) ResetIncrement() error {
	args := m.Called()
	return args.Error(0)
}
