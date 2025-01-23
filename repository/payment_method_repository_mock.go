package repository

import (
	"clothing-pair-project/entity"

	"github.com/stretchr/testify/mock"
)

type PaymentMethodRepositoryMock struct {
	mock.Mock
}

func (m *PaymentMethodRepositoryMock) GetAll() ([]entity.PaymentMethod, error) {
	args := m.Called()
	return args.Get(0).([]entity.PaymentMethod), args.Error(1)
}

func (m *PaymentMethodRepositoryMock) GetByID(paymentMethodID int) (entity.PaymentMethod, error) {
	args := m.Called(paymentMethodID)
	return args.Get(0).(entity.PaymentMethod), args.Error(1)
}
