package repository

import (
	"clothing-pair-project/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(id int) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) FindByName(name string) ([]models.Product, error) {
	args := m.Called(name)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) FindByCategoryID(categoryID int) ([]models.Product, error) {
	args := m.Called(categoryID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) Add(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
