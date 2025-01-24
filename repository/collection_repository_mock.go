package repository

import (
	"clothing-pair-project/entity"

	"github.com/stretchr/testify/mock"
)

type CollectionRepositoryMock struct {
	mock.Mock
}

func (m *CollectionRepositoryMock) FindAll() ([]entity.Collection, error) {
	args := m.Called()
	return args.Get(0).([]entity.Collection), args.Error(1)
}

func (m *CollectionRepositoryMock) FindByID(collectionID int) (entity.Collection, error) {
	args := m.Called(collectionID)
	return args.Get(0).(entity.Collection), args.Error(1)
}

func (m *CollectionRepositoryMock) Add(collection entity.Collection) error {
	args := m.Called(collection)
	return args.Error(0)
}

func (m *CollectionRepositoryMock) Update(collection entity.Collection) error {
	args := m.Called(collection)
	return args.Error(0)
}

func (m *CollectionRepositoryMock) Delete(collectionID int) error {
	args := m.Called(collectionID)
	return args.Error(0)
}

func (m *CollectionRepositoryMock) ResetIncrement() error {
	args := m.Called()
	return args.Error(0)
}
