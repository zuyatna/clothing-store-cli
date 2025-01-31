package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var productDetailRepository = repository.MockProductDetailRepository{Mock: mock.Mock{}}
var productDetailService = services.NewProductDetailService(&productDetailRepository)

func TestGetAllProductDetails(t *testing.T) {
	productDetails := []models.ProductDetail{
		{
			ProductDetailID: 1,
			ProductID:       1,
			SizeID:          1,
			Stock:           10,
		},
		{
			ProductDetailID: 2,
			ProductID:       2,
			SizeID:          2,
			Stock:           10,
		},
	}

	productDetailRepository.On("FindAll").Return(productDetails, nil)

	result, err := productDetailService.GetAllProductDetails()
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}

	if len(result) != len(productDetails) {
		t.Errorf("Length of result is not equal to length of productDetails")
	}

	for i := range result {
		if result[i] != productDetails[i] {
			t.Errorf("Result is not equal to productDetails")
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, productDetails, result)

	productDetailRepository.AssertExpectations(t)
}

func TestGetProductDetailByID(t *testing.T) {
	t.Run("Success Get Product Detail By ID", func(t *testing.T) {
		productDetail := models.ProductDetail{
			ProductDetailID: 1,
			ProductID:       1,
			SizeID:          1,
			Stock:           10,
		}

		productDetailRepository.On("FindByID", 1).Return(productDetail, nil)

		result, err := productDetailService.GetProductDetailByID(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, productDetail, result)

		productDetailRepository.AssertExpectations(t)
	})

	t.Run("Failed Get Product Detail By ID", func(t *testing.T) {
		productDetailRepository.On("FindByID", 2).Return(models.ProductDetail{}, assert.AnError)

		result, err := productDetailService.GetProductDetailByID(2)
		if err == nil {
			t.Errorf("Error was expected: %s", err)
		}

		assert.Error(t, err)
		assert.Equal(t, models.ProductDetail{}, result)

		productDetailRepository.AssertExpectations(t)
	})
}
