package repository

import (
	"clothing-pair-project/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockProductDetailRepository struct {
	mock.Mock
}

func (m *MockProductDetailRepository) FindAll() ([]models.ProductDetail, error) {
	args := m.Called()
	return args.Get(0).([]models.ProductDetail), args.Error(1)
}

func (m *MockProductDetailRepository) FindByID(id int) (models.ProductDetail, error) {
	args := m.Called(id)
	return args.Get(0).(models.ProductDetail), args.Error(1)
}

func (m *MockProductDetailRepository) Add(productDetail models.ProductDetail) error {
	args := m.Called(productDetail)
	return args.Error(0)
}

func (m *MockProductDetailRepository) Update(productDetail models.ProductDetail) error {
	args := m.Called(productDetail)
	return args.Error(0)
}

func (m *MockProductDetailRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
