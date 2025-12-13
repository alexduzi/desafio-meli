package usecase

import (
	"fmt"
	"testing"
	"time"

	"project/internal/entity"
	"project/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestListProductUseCase_Execute_Success(t *testing.T) {
	now := time.Now()
	mockRepo := &MockProductRepository{
		ListProductsFunc: func() ([]entity.Product, error) {
			return []entity.Product{
				{
					ID:          "PROD-1",
					Title:       "Product 1",
					Description: "Description 1",
					Price:       100.0,
					Currency:    "USD",
					Condition:   "new",
					Stock:       10,
					SellerID:    "seller-1",
					SellerName:  "Store 1",
					Category:    "Electronics",
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				{
					ID:          "PROD-2",
					Title:       "Product 2",
					Description: "Description 2",
					Price:       200.0,
					Currency:    "USD",
					Condition:   "used",
					Stock:       5,
					SellerID:    "seller-2",
					SellerName:  "Store 2",
					Category:    "Books",
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			}, nil
		},
	}

	useCase := NewListProductUseCase(mockRepo)
	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "PROD-1", result[0].ID)
	assert.Equal(t, "Product 1", result[0].Title)
	assert.Equal(t, 100.0, result[0].Price)
	assert.Equal(t, "PROD-2", result[1].ID)
}

func TestListProductUseCase_Execute_EmptyList(t *testing.T) {
	mockRepo := &MockProductRepository{
		ListProductsFunc: func() ([]entity.Product, error) {
			return []entity.Product{}, nil
		},
	}

	useCase := NewListProductUseCase(mockRepo)
	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result)
}

func TestListProductUseCase_Execute_DatabaseError(t *testing.T) {
	mockRepo := &MockProductRepository{
		ListProductsFunc: func() ([]entity.Product, error) {
			return nil, fmt.Errorf("%w: connection timeout", errors.ErrDatabaseError)
		},
	}

	useCase := NewListProductUseCase(mockRepo)
	result, err := useCase.Execute()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.ErrDatabaseError)
}

func TestListProductUseCase_Execute_MapsAllFields(t *testing.T) {
	now := time.Now()
	mockRepo := &MockProductRepository{
		ListProductsFunc: func() ([]entity.Product, error) {
			return []entity.Product{
				{
					ID:          "PROD-TEST",
					Title:       "Test Title",
					Description: "Test Description",
					Price:       99.99,
					Currency:    "EUR",
					Condition:   "refurbished",
					Stock:       3,
					SellerID:    "seller-test",
					SellerName:  "Test Seller",
					Category:    "Test Category",
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			}, nil
		},
	}

	useCase := NewListProductUseCase(mockRepo)
	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Len(t, result, 1)

	product := result[0]
	assert.Equal(t, "PROD-TEST", product.ID)
	assert.Equal(t, "Test Title", product.Title)
	assert.Equal(t, "Test Description", product.Description)
	assert.Equal(t, 99.99, product.Price)
	assert.Equal(t, "EUR", product.Currency)
	assert.Equal(t, "refurbished", product.Condition)
	assert.Equal(t, 3, product.Stock)
	assert.Equal(t, "seller-test", product.SellerID)
	assert.Equal(t, "Test Seller", product.SellerName)
	assert.Equal(t, "Test Category", product.Category)
	assert.Equal(t, now, product.CreatedAt)
	assert.Equal(t, now, product.UpdatedAt)
}
