package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/services"
	"clothing-pair-project/tests/unit/repository"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = repository.MockProductRepository{Mock: mock.Mock{}}
var productService = services.NewProductService(&productRepository)

func TestGetAllProducts(t *testing.T) {
	products := []models.Product{
		{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type: "men",
		},
		{
			ProductID:  2,
			CategoryID: 1,
			Name:       "product2",
			Price:      249000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type: "women",
		},
	}

	productRepository.On("FindAll", 5, 0).Return(products, nil)

	result, err := productService.GetAllProducts(5, 0)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, products, result)

	productRepository.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	t.Run("Success Get Product By ID", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("FindByID", 1).Return(product, nil)

		result, err := productService.GetProductByID(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, product, result)

		productRepository.AssertExpectations(t)
	})

	t.Run("Product Not Found", func(t *testing.T) {
		productRepository.On("FindByID", 2).Return(models.Product{}, errors.New("product not found"))

		result, err := productService.GetProductByID(2)
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)
		assert.Equal(t, models.Product{}, result)

		productRepository.AssertExpectations(t)
	})
}

func TestGetProductByName(t *testing.T) {
	t.Run("Success Get Product By Name", func(t *testing.T) {
		products := []models.Product{
			{
				ProductID:  1,
				CategoryID: 1,
				Name:       "product1",
				Price:      149000,
				Description: sql.NullString{
					String: "description1",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "image1",
					Valid:  true,
				},
				Type: "men",
			},
		}

		productRepository.On("FindByName", "product1").Return(products, nil)

		result, err := productService.GetProductByName("product1")
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, products, result)

		productRepository.AssertExpectations(t)
	})

	t.Run("Product Not Found", func(t *testing.T) {
		productRepository.On("FindByName", "product2").Return([]models.Product{}, errors.New("product not found"))

		result, err := productService.GetProductByName("product2")
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)
		assert.Equal(t, []models.Product{}, result)

		productRepository.AssertExpectations(t)
	})
}

func TestGetProductByCategoryID(t *testing.T) {
	t.Run("Success Get Product By Category ID", func(t *testing.T) {
		products := []models.Product{
			{
				ProductID:  1,
				CategoryID: 1,
				Name:       "product1",
				Price:      149000,
				Description: sql.NullString{
					String: "description1",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "image1",
					Valid:  true,
				},
				Type:      "men",
				CreatedAt: time.Time{},
			},
			{
				ProductID:  2,
				CategoryID: 1,
				Name:       "product2",
				Price:      249000,
				Description: sql.NullString{
					String: "description1",
					Valid:  true,
				},
				Images: sql.NullString{
					String: "image1",
					Valid:  true,
				},
				Type:      "women",
				CreatedAt: time.Time{},
			},
		}

		productRepository.On("FindByCategoryID", 1).Return(products, nil)

		result, err := productService.GetProductByCategoryID(1, 5, 0)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, products, result)

		productRepository.AssertExpectations(t)
	})

	t.Run("Category Not Found", func(t *testing.T) {
		productRepository.On("FindByCategoryID", 2).Return([]models.Product{}, errors.New("category not found"))

		result, err := productService.GetProductByCategoryID(2, 5, 0)
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)
		assert.Equal(t, []models.Product{}, result)

		productRepository.AssertExpectations(t)
	})
}

func TestAddProduct(t *testing.T) {
	t.Run("Success Add Product", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(nil)

		err := productService.AddProduct(product)
		assert.NoError(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Category Not Found", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 2,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(errors.New("category not found"))

		err := productService.AddProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Type", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(errors.New("invalid type"))

		err := productService.AddProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Price", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      -149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(errors.New("invalid price"))

		err := productService.AddProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Name", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(errors.New("invalid name"))

		err := productService.AddProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Failed to Add Product", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "",
			Price:      0,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "",
			CreatedAt: time.Time{},
		}

		productRepository.On("Add", product).Return(errors.New("failed to add product"))

		err := productService.AddProduct(product)
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)

		productRepository.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Success Update Product", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(nil)

		err := productService.UpdateProduct(product)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, nil, err)

		productRepository.AssertExpectations(t)
	})

	t.Run("Category Not Found", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 2,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(errors.New("category not found"))

		err := productService.UpdateProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Type", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(errors.New("invalid type"))

		err := productService.UpdateProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Price", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "product1",
			Price:      -149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(errors.New("invalid price"))

		err := productService.UpdateProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Invalid Name", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "",
			Price:      149000,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "men",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(errors.New("invalid name"))

		err := productService.UpdateProduct(product)
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("Failed to Update Product", func(t *testing.T) {
		product := models.Product{
			ProductID:  1,
			CategoryID: 1,
			Name:       "",
			Price:      0,
			Description: sql.NullString{
				String: "description1",
				Valid:  true,
			},
			Images: sql.NullString{
				String: "image1",
				Valid:  true,
			},
			Type:      "",
			CreatedAt: time.Time{},
		}

		productRepository.On("Update", product).Return(errors.New("failed to update product"))

		err := productService.UpdateProduct(product)
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)

		productRepository.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Success Delete Product", func(t *testing.T) {
		productRepository.On("Delete", 1).Return(nil)

		err := productService.DeleteProduct(1)
		if err != nil {
			t.Errorf("Error was not expected: %s", err)
		}

		assert.NoError(t, err)

		productRepository.AssertExpectations(t)
	})

	t.Run("Product Not Found", func(t *testing.T) {
		productRepository.On("Delete", 2).Return(errors.New("product not found"))

		err := productService.DeleteProduct(2)
		if err == nil {
			t.Errorf("Error was expected, got nil")
		}

		assert.Error(t, err)

		productRepository.AssertExpectations(t)
	})
}
