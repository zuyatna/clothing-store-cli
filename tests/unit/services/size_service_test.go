package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var sizeRepository = repository.MockSizeRepository{Mock: mock.Mock{}}
var sizeService = services.NewSizeService(&sizeRepository)

func TestGetAllSizes(t *testing.T) {
	sizeRepository.On("FindAll").Return([]models.Size{}, nil)
	sizes, err := sizeService.GetAllSizes()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(sizes))
	sizeRepository.AssertExpectations(t)
}
