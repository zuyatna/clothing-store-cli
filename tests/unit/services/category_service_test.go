package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = repository.MockCategoryRepository{Mock: mock.Mock{}}
var categoryService = services.NewCategoryService(&categoryRepository)

func TestGetAllCategories(t *testing.T) {
	categories := []models.Category{
		{
			CategoryID: 1,
			Name:       "category1",
		},
		{
			CategoryID: 2,
			Name:       "category2",
		},
	}

	categoryRepository.On("FindAll").Return(categories, nil)

	result, err := categoryService.GetAllCategories()
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}

	if len(result) == 0 {
		t.Error("Result is empty")
	}

	assert.NoError(t, err)
	assert.Equal(t, categories, result)

	categoryRepository.AssertExpectations(t)
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("Success Get Category By ID", func(t *testing.T) {
		category := models.Category{
			CategoryID: 1,
			Name:       "category1",
		}

		categoryRepository.On("FindByID", 1).Return(category, nil)

		result, err := categoryService.GetCategoryByID(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, category, result)

		categoryRepository.AssertExpectations(t)
	})

	t.Run("Failed Get Category By ID", func(t *testing.T) {
		category := models.Category{}

		categoryRepository.On("FindByID", 2).Return(category, assert.AnError)

		result, err := categoryService.GetCategoryByID(2)
		if err == nil {
			t.Error("Error was expected")
		}

		assert.Error(t, err)
		assert.Equal(t, category, result)

		categoryRepository.AssertExpectations(t)
	})
}

func TestAddCategory(t *testing.T) {
	t.Run("Success Add Category", func(t *testing.T) {
		category := models.Category{
			CategoryID: 1,
			Name:       "category1",
		}

		categoryRepository.On("Add", category).Return(nil)

		err := categoryService.AddCategory(category)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, category.Name, "category1")

		categoryRepository.AssertExpectations(t)
	})

	t.Run("Failed Add Category", func(t *testing.T) {
		category := models.Category{}

		categoryRepository.On("Add", category).Return(assert.AnError)

		err := categoryService.AddCategory(category)
		if err == nil {
			t.Error("Error was expected")
		}

		assert.Error(t, err)

		categoryRepository.AssertExpectations(t)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Success Update Category", func(t *testing.T) {
		category := models.Category{
			CategoryID: 1,
			Name:       "category1",
		}

		categoryRepository.On("Update", category).Return(nil)

		err := categoryService.UpdateCategory(category)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, category.Name, "category1")

		categoryRepository.AssertExpectations(t)
	})

	t.Run("Failed Update Category", func(t *testing.T) {
		category := models.Category{}

		categoryRepository.On("Update", category).Return(assert.AnError)

		err := categoryService.UpdateCategory(category)
		if err == nil {
			t.Error("Error was expected")
		}

		assert.Error(t, err)

		categoryRepository.AssertExpectations(t)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("Success Delete Category", func(t *testing.T) {
		categoryRepository.On("Delete", 1).Return(nil)

		err := categoryService.DeleteCategory(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)

		categoryRepository.AssertExpectations(t)
	})

	t.Run("Failed Delete Category", func(t *testing.T) {
		categoryRepository.On("Delete", 2).Return(assert.AnError)

		err := categoryService.DeleteCategory(2)
		if err == nil {
			t.Error("Error was expected")
		}

		assert.Error(t, err)

		categoryRepository.AssertExpectations(t)
	})
}
