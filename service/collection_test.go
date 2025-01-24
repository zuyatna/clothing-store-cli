package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var collectionServiceRepository = &repository.CollectionRepositoryMock{Mock: mock.Mock{}}
var collectionService = NewCollectionService(collectionServiceRepository)

func TestFindAllCollection(t *testing.T) {
	collections := []entity.Collection{
		{
			CollectionID: 1,
			Name:         "T-Shirt",
			CreatedAt:    time.Now(),
		},
		{
			CollectionID: 2,
			Name:         "Pants",
			CreatedAt:    time.Now(),
		},
	}

	collectionServiceRepository.On("FindAll").Return(collections, nil)

	result, err := collectionService.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, collections, result)
	collectionServiceRepository.AssertExpectations(t)
}

func TestFindCollectionByID(t *testing.T) {
	collection := entity.Collection{
		CollectionID: 1,
		Name:         "T-Shirt",
		CreatedAt:    time.Now(),
	}

	collectionServiceRepository.On("FindByID", 1).Return(collection, nil)

	result, err := collectionService.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, collection, result)
	collectionServiceRepository.AssertExpectations(t)
}

func TestAddCollection(t *testing.T) {
	t.Run("add collection success", func(t *testing.T) {
		collection := entity.Collection{
			CollectionID: 1,
			Name:         "T-Shirt",
		}

		collectionServiceRepository.On("Add", collection).Return(nil).Once()

		err := collectionService.Add(collection)
		assert.NoError(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})

	t.Run("add collection failed", func(t *testing.T) {
		collection := entity.Collection{
			CollectionID: 1,
			Name:         "T-Shirt",
		}

		collectionServiceRepository.On("Add", collection).Return(errDummy).Once()

		err := collectionService.Add(collection)
		assert.Error(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})
}

func TestUpdateCollection(t *testing.T) {
	t.Run("update collection success", func(t *testing.T) {
		collection := entity.Collection{
			CollectionID: 1,
			Name:         "T-Shirt",
		}

		collectionServiceRepository.On("Update", collection).Return(nil).Once()

		err := collectionService.Update(collection)
		assert.NoError(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})

	t.Run("update collection failed", func(t *testing.T) {
		collection := entity.Collection{
			CollectionID: 1,
			Name:         "T-Shirt",
		}

		collectionServiceRepository.On("Update", collection).Return(errDummy).Once()

		err := collectionService.Update(collection)
		assert.Error(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})
}

func TestDeleteCollection(t *testing.T) {
	t.Run("delete collection success", func(t *testing.T) {
		collectionID := 1

		collectionServiceRepository.On("Delete", collectionID).Return(nil).Once()

		err := collectionService.Delete(collectionID)
		assert.NoError(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})

	t.Run("delete collection failed", func(t *testing.T) {
		collectionID := 1

		collectionServiceRepository.On("Delete", collectionID).Return(errDummy).Once()

		err := collectionService.Delete(collectionID)
		assert.Error(t, err)
		collectionServiceRepository.AssertExpectations(t)
	})
}

func TestResetIncrementCollection(t *testing.T) {
	collectionServiceRepository.On("ResetIncrement").Return(nil).Once()

	err := collectionService.ResetIncrement()
	assert.NoError(t, err)
	collectionServiceRepository.AssertExpectations(t)
}
