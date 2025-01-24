package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userServiceRepository = &repository.UsersRepositoryMock{Mock: mock.Mock{}}
var userService = NewUserService(userServiceRepository)

func TestFindAllUser(t *testing.T) {
	users := []entity.User{
		{
			UserID:   1,
			Username: "user1",
			Email:    "user1@email.com",
			Password: "password1",
			Role:     "user",
		},
		{
			UserID:   2,
			Username: "user2",
			Email:    "user2@email.com",
			Password: "password2",
			Role:     "user",
		},
	}

	userServiceRepository.On("FindAll").Return(users, nil)

	result, err := userService.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, users, result)
	userServiceRepository.AssertExpectations(t)
}

func TestFindUserByID(t *testing.T) {
	user := entity.User{
		UserID:   1,
		Username: "user1",
		Email:    "user1@email.com",
		Password: "password1",
		Role:     "user",
	}

	userServiceRepository.On("FindByID", 1).Return(user, nil)

	result, err := userService.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	userServiceRepository.AssertExpectations(t)
}

func TestFindUserByUsername(t *testing.T) {
	user := entity.User{
		UserID:   1,
		Username: "user1",
		Email:    "user1@email.com",
		Password: "password1",
		Role:     "user",
	}

	userServiceRepository.On("FindByUsername", "user1").Return(user, nil)

	result, err := userService.FindByUsername("user1")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	userServiceRepository.AssertExpectations(t)
}

func TestAddUser(t *testing.T) {
	t.Run("add user success", func(t *testing.T) {
		user := entity.User{
			UserID:   1,
			Username: "user1",
			Email:    "user1@email.com",
			Password: "password1",
			Role:     "user",
		}

		userServiceRepository.On("Add", user).Return(nil).Once()

		err := userService.Add(user)
		assert.NoError(t, err)
		assert.Contains(t, user.Email, "@", "Email should contain @")
		userServiceRepository.AssertExpectations(t)
	})

	t.Run("add user with invalid email", func(t *testing.T) {
		user := entity.User{
			UserID:   1,
			Username: "user1",
			Email:    "invalidemail", // Email without @
			Password: "password1",
			Role:     "user",
		}

		userServiceRepository.On("Add", user).Return(errDummy).Once()

		err := userService.Add(user)
		assert.Error(t, err)
		assert.NotContains(t, user.Email, "@", "Email should contain @")
		userServiceRepository.AssertExpectations(t)
	})

	t.Run("add user failed", func(t *testing.T) {
		user := entity.User{
			UserID:   1,
			Username: "user1",
			Email:    "user1@email.com",
			Password: "password1",
			Role:     "user",
		}

		userServiceRepository.On("Add", user).Return(errDummy).Once()

		err := userService.Add(user)
		assert.Error(t, err)
		userServiceRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("update user success", func(t *testing.T) {
		user := entity.User{
			UserID:   1,
			Username: "user1",
			Email:    "user12@email.com",
			Password: "password12",
			Role:     "user",
		}

		userServiceRepository.On("Update", user).Return(nil).Once()

		err := userService.Update(user)
		assert.NoError(t, err)
		userServiceRepository.AssertExpectations(t)
	})

	t.Run("update user failed", func(t *testing.T) {
		user := entity.User{
			UserID:   1,
			Username: "user12",
			Email:    "user1@email.com",
			Password: "password12",
			Role:     "user",
		}

		userServiceRepository.On("Update", user).Return(errDummy).Once()

		err := userService.Update(user)
		assert.Error(t, err)
		userServiceRepository.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("delete user success", func(t *testing.T) {
		userServiceRepository.On("Delete", 1).Return(nil).Once()

		err := userService.Delete(1)
		assert.NoError(t, err)
		userServiceRepository.AssertExpectations(t)
	})

	t.Run("delete user failed", func(t *testing.T) {
		userServiceRepository.On("Delete", 1).Return(errDummy).Once()

		err := userService.Delete(1)
		assert.Error(t, err)
		userServiceRepository.AssertExpectations(t)
	})
}

func TestResetIncrementUser(t *testing.T) {
	userServiceRepository.On("ResetIncrement").Return(nil)

	err := userService.ResetIncrement()
	assert.NoError(t, err)
	userServiceRepository.AssertExpectations(t)
}
