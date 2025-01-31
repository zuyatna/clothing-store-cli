package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var colorRepository = repository.MockColorRepository{Mock: mock.Mock{}}
var colorService = services.NewColorService(&colorRepository)

func TestGetAllColors(t *testing.T) {
	colorRepository.On("FindAll").Return([]models.Color{}, nil)
	colors, err := colorService.GetAllColors()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(colors))
	colorRepository.AssertExpectations(t)
}
