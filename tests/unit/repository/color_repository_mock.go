package repository

import (
	"clothing-pair-project/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockColorRepository struct {
	mock.Mock
}

func (m *MockColorRepository) FindAll() ([]models.Color, error) {
	args := m.Called()
	return args.Get(0).([]models.Color), args.Error(1)
}

func (m *MockColorRepository) FindByID(id int) (models.Color, error) {
	args := m.Called(id)
	return args.Get(0).(models.Color), args.Error(1)
}

func (m *MockColorRepository) Add(color models.Color) error {
	args := m.Called(color)
	return args.Error(0)
}

func (m *MockColorRepository) Update(color models.Color) error {
	args := m.Called(color)
	return args.Error(0)
}

func (m *MockColorRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
