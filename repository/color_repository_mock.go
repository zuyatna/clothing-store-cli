package repository

import (
	"clothing-pair-project/entity"

	"github.com/stretchr/testify/mock"
)

type ColorRepositoryMock struct {
	mock.Mock
}

func (m *ColorRepositoryMock) FindAll() ([]entity.Color, error) {
	args := m.Called()
	return args.Get(0).([]entity.Color), args.Error(1)
}

func (m *ColorRepositoryMock) FindByID(colorID int) (entity.Color, error) {
	args := m.Called(colorID)
	return args.Get(0).(entity.Color), args.Error(1)
}

func (m *ColorRepositoryMock) Add(color entity.Color) error {
	args := m.Called(color)
	return args.Error(0)
}

func (m *ColorRepositoryMock) Update(color entity.Color) error {
	args := m.Called(color)
	return args.Error(0)
}

func (m *ColorRepositoryMock) Delete(colorID int) error {
	args := m.Called(colorID)
	return args.Error(0)
}

func (m *ColorRepositoryMock) ResetIncrement() error {
	args := m.Called()
	return args.Error(0)
}
