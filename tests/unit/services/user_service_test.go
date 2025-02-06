package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = repository.MockUserRepository{Mock: mock.Mock{}}
var userService = services.NewUserService(&userRepository)

func TestGetAllUsers(t *testing.T) {
	users := []models.User{
		{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		},
		{
			UserID:    2,
			Username:  "user2",
			Email:     "user2@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		},
	}

	userRepository.On("FindAll", 5, 0).Return(users, nil)

	result, err := userService.GetAllUsers(5, 0)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}

	if len(result) != len(users) {
		t.Errorf("Length of result is not equal to length of users")
	}

	for i := range result {
		if result[i] != users[i] {
			t.Errorf("Result is not equal to users")
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, users, result)

	userRepository.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	t.Run("Success Get User By ID", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("FindByID", 1).Return(user, nil)

		result, err := userService.GetUserByID(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, user, result)

		userRepository.AssertExpectations(t)
	})

	t.Run("Failed Get User By ID", func(t *testing.T) {
		userRepository.On("FindByID", 2).Return(models.User{}, errors.New("user not found"))

		result, err := userService.GetUserByID(2)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Equal(t, models.User{}, result)

		userRepository.AssertExpectations(t)
	})
}

func TestGetUserByUsername(t *testing.T) {
	t.Run("Success Get User By Username", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("FindByUsername", "user1").Return(user, nil)

		result, err := userService.GetUserByUsername("user1")
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, user, result)

		userRepository.AssertExpectations(t)
	})

	t.Run("Failed Get User By Username", func(t *testing.T) {
		userRepository.On("FindByUsername", "user2").Return(models.User{}, errors.New("user not found"))

		result, err := userService.GetUserByUsername("user2")
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Equal(t, models.User{}, result)

		userRepository.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {
	t.Run("Success Add User", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Add", user).Return(nil)

		err := userService.AddUser(user)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Contains(t, user.Email, "@", "email should contain @")
		assert.NotContains(t, user.Username, " ", "username should not contain spaces")

		userRepository.AssertExpectations(t)
	})

	t.Run("Add User With Invalid Username", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user 1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Add", user).Return(errors.New("username should not contain spaces"))

		err := userService.AddUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username should not contain spaces")

		userRepository.AssertExpectations(t)
	})

	t.Run("Add User With Invalid Email", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Add", user).Return(errors.New("email should contain @"))

		err := userService.AddUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email should contain @")

		userRepository.AssertExpectations(t)
	})

	t.Run("Add User With Invalid Role", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1email.com",
			Password:  "password",
			Role:      "superuser",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Add", user).Return(errors.New("role should be either user or admin"))

		err := userService.AddUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "role should be either user or admin")

		userRepository.AssertExpectations(t)
	})

	t.Run("Failed Add User", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user 1",
			Email:     "user1email.com",
			Password:  "password",
			Role:      "",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Add", user).Return(errors.New("failed to add user"))

		err := userService.AddUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Success Update User", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Update", user).Return(nil)

		err := userService.UpdateUser(user)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Contains(t, user.Email, "@", "email should contain @")
		assert.NotContains(t, user.Username, " ", "username should not contain spaces")

		userRepository.AssertExpectations(t)
	})

	t.Run("Update User With Invalid Username", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user 1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Update", user).Return(errors.New("username should not contain spaces"))

		err := userService.UpdateUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username should not contain spaces")

		userRepository.AssertExpectations(t)
	})

	t.Run("Update User With Invalid Email", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Update", user).Return(errors.New("email should contain @"))

		err := userService.UpdateUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email should contain @")

		userRepository.AssertExpectations(t)
	})

	t.Run("Update User With Invalid Role", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user1",
			Email:     "user1@email.com",
			Password:  "password",
			Role:      "superuser",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Update", user).Return(errors.New("role should be either user or admin"))

		err := userService.UpdateUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "role should be either user or admin")

		userRepository.AssertExpectations(t)
	})

	t.Run("Failed Update User", func(t *testing.T) {
		user := models.User{
			UserID:    1,
			Username:  "user 1",
			Email:     "user1email.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			Active:    true,
		}

		userRepository.On("Update", user).Return(errors.New("failed to update user"))

		err := userService.UpdateUser(user)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Success Delete User", func(t *testing.T) {
		userRepository.On("Delete", 1).Return(nil)

		err := userService.DeleteUser(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Failed Delete User", func(t *testing.T) {
		userRepository.On("Delete", 2).Return(errors.New("failed to delete user"))

		err := userService.DeleteUser(2)
		if err == nil {
			t.Errorf("Error was expected")
		}

		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})
}
