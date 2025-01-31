package repository

import (
	"clothing-pair-project/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockSizeRepository struct {
	mock.Mock
}

func (m *MockSizeRepository) FindAll() ([]models.Size, error) {
	args := m.Called()
	return args.Get(0).([]models.Size), args.Error(1)
}

func (m *MockSizeRepository) FindByID(id int) (models.Size, error) {
	args := m.Called(id)
	return args.Get(0).(models.Size), args.Error(1)
}

func (m *MockSizeRepository) Add(size models.Size) error {
	args := m.Called(size)
	return args.Error(0)
}

func (m *MockSizeRepository) Update(size models.Size) error {
	args := m.Called(size)
	return args.Error(0)
}

func (m *MockSizeRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
