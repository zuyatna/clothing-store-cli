package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var colorServiceRepository = &repository.ColorRepositoryMock{Mock: mock.Mock{}}
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
	color := entity.Color{
		ColorID: 1,
		Name:    "Red",
	}

	colorServiceRepository.On("Add", color).Return(nil)

	err := colorService.Add(color)
	assert.NoError(t, err)
	colorServiceRepository.AssertExpectations(t)
}

func TestUpdateColor(t *testing.T) {
	color := entity.Color{
		ColorID: 1,
		Name:    "Red",
	}

	colorServiceRepository.On("Update", color).Return(nil)

	err := colorService.Update(color)
	assert.NoError(t, err)
	colorServiceRepository.AssertExpectations(t)
}

func TestDeleteColor(t *testing.T) {
	colorServiceRepository.On("Delete", 1).Return(nil)

	err := colorService.Delete(1)
	assert.NoError(t, err)
	colorServiceRepository.AssertExpectations(t)
}

func TestResetIncrementColor(t *testing.T) {
	colorServiceRepository.On("ResetIncrement").Return(nil)

	err := colorService.ResetIncrement()
	assert.NoError(t, err)
	colorServiceRepository.AssertExpectations(t)
}
