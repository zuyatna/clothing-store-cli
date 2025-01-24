package repository

import (
	"clothing-pair-project/entity"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func (m *UsersRepositoryMock) FindAll() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *UsersRepositoryMock) FindByID(userID int) (entity.User, error) {
	args := m.Called(userID)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UsersRepositoryMock) FindByUsername(username string) (entity.User, error) {
	args := m.Called(username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UsersRepositoryMock) Add(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UsersRepositoryMock) Update(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UsersRepositoryMock) Delete(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *UsersRepositoryMock) ResetIncrement() error {
	args := m.Called()
	return args.Error(0)
}
