package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var colorServiceRepository = &repository.ColorRepositoryMock{Mock: mock.Mock{}}
var errDummy = errors.New("dummy error")
var colorService = NewColorService(colorServiceRepository)

func TestFindAllColor(t *testing.T) {
	colors := []entity.Color{
		{
			ColorID: 1,
			Name:    "Red",
		},
		{
			ColorID: 2,
			Name:    "Blue",
		},
	}

	colorServiceRepository.On("FindAll").Return(colors, nil)

	result, err := colorService.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, colors, result)
	colorServiceRepository.AssertExpectations(t)
}

func TestFindColorByID(t *testing.T) {
	color := entity.Color{
		ColorID: 1,
		Name:    "Red",
	}

	colorServiceRepository.On("FindByID", 1).Return(color, nil)

	result, err := colorService.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, color, result)
	colorServiceRepository.AssertExpectations(t)
}

func TestAddColor(t *testing.T) {
	t.Run("add color success", func(t *testing.T) {
		color := entity.Color{
			ColorID: 1,
			Name:    "Red",
		}

		colorServiceRepository.On("Add", color).Return(nil).Once()

		err := colorService.Add(color)
		assert.NoError(t, err)
		colorServiceRepository.AssertExpectations(t)
	})

	t.Run("add color failed", func(t *testing.T) {
		color := entity.Color{
			ColorID: 1,
			Name:    "Red",
		}

		colorServiceRepository.On("Add", color).Return(errDummy).Once()

		err := colorService.Add(color)
		assert.Error(t, err)
		colorServiceRepository.AssertExpectations(t)
	})
}

func TestUpdateColor(t *testing.T) {
	t.Run("update color success", func(t *testing.T) {
		color := entity.Color{
			ColorID: 1,
			Name:    "Red",
		}

		colorServiceRepository.On("Update", color).Return(nil).Once()

		err := colorService.Update(color)
		assert.NoError(t, err)
		colorServiceRepository.AssertExpectations(t)
	})

	t.Run("update color failed", func(t *testing.T) {
		color := entity.Color{
			ColorID: 1,
			Name:    "Red",
		}

		colorServiceRepository.On("Update", color).Return(errDummy).Once()

		err := colorService.Update(color)
		assert.Error(t, err)
		colorServiceRepository.AssertExpectations(t)
	})
}

func TestDeleteColor(t *testing.T) {
	t.Run("delete color success", func(t *testing.T) {
		colorServiceRepository.On("Delete", 1).Return(nil).Once()

		err := colorService.Delete(1)
		assert.NoError(t, err)
		colorServiceRepository.AssertExpectations(t)
	})

	t.Run("delete color failed", func(t *testing.T) {
		colorServiceRepository.On("Delete", 1).Return(errDummy).Once()

		err := colorService.Delete(1)
		assert.Error(t, err)
		colorServiceRepository.AssertExpectations(t)
	})
}

func TestResetIncrementColor(t *testing.T) {
	colorServiceRepository.On("ResetIncrement").Return(nil)

	err := colorService.ResetIncrement()
	assert.NoError(t, err)
	colorServiceRepository.AssertExpectations(t)
}
