package usecase

import (
	"fmt"
	"testing"
	"time"

	"project/internal/entity"
	"project/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestGetProductUseCase_Execute_Success(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetProductFunc: func(id string) (*entity.Product, error) {
			return &entity.Product{
				ID:          "PROD-123",
				Title:       "iPhone 15",
				Description: "Latest iPhone",
				Price:       999.99,
				Currency:    "USD",
				Condition:   "new",
				Stock:       10,
				SellerID:    "seller-1",
				SellerName:  "Apple Store",
				Category:    "Electronics",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, nil
		},
		FindImagesByProductIDFunc: func(productID string) ([]entity.ProductImage, error) {
			return []entity.ProductImage{
				{ID: 1, ProductID: "PROD-123", ImageURL: "http://example.com/image1.jpg", DisplayOrder: 0},
				{ID: 2, ProductID: "PROD-123", ImageURL: "http://example.com/image2.jpg", DisplayOrder: 1},
			}, nil
		},
	}

	useCase := NewGetProductUseCase(mockRepo)
	result, err := useCase.Execute(ProductInputDTO{ID: "PROD-123"})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "PROD-123", result.ID)
	assert.Equal(t, "iPhone 15", result.Title)
	assert.Equal(t, 999.99, result.Price)
	assert.Len(t, result.Images, 2)
	assert.Equal(t, "http://example.com/image1.jpg", result.Images[0].ImageURL)
}

func TestGetProductUseCase_Execute_EmptyID(t *testing.T) {
	mockRepo := &MockProductRepository{}
	useCase := NewGetProductUseCase(mockRepo)

	tests := []struct {
		name string
		id   string
	}{
		{"Empty string", ""},
		{"Only spaces", "   "},
		{"Only tabs", "\t\t"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := useCase.Execute(ProductInputDTO{ID: tt.id})

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Equal(t, errors.ErrInvalidProductID, err)
		})
	}
}

func TestGetProductUseCase_Execute_ProductNotFound(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetProductFunc: func(id string) (*entity.Product, error) {
			return nil, errors.ErrProductNotFound
		},
	}

	useCase := NewGetProductUseCase(mockRepo)
	result, err := useCase.Execute(ProductInputDTO{ID: "PROD-999"})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.ErrProductNotFound)
}

func TestGetProductUseCase_Execute_DatabaseError(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetProductFunc: func(id string) (*entity.Product, error) {
			return nil, fmt.Errorf("%w: connection failed", errors.ErrDatabaseError)
		},
	}

	useCase := NewGetProductUseCase(mockRepo)
	result, err := useCase.Execute(ProductInputDTO{ID: "PROD-123"})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.ErrDatabaseError)
}

func TestGetProductUseCase_Execute_ImagesError(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetProductFunc: func(id string) (*entity.Product, error) {
			return &entity.Product{
				ID:    "PROD-123",
				Title: "Test Product",
			}, nil
		},
		FindImagesByProductIDFunc: func(productID string) ([]entity.ProductImage, error) {
			return nil, fmt.Errorf("%w: failed to fetch images", errors.ErrDatabaseError)
		},
	}

	useCase := NewGetProductUseCase(mockRepo)
	result, err := useCase.Execute(ProductInputDTO{ID: "PROD-123"})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.ErrDatabaseError)
}

func TestGetProductUseCase_Execute_NoImages(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetProductFunc: func(id string) (*entity.Product, error) {
			return &entity.Product{
				ID:          "PROD-123",
				Title:       "Product Without Images",
				Description: "Test",
				Price:       50.0,
				Currency:    "USD",
				Condition:   "new",
				Stock:       5,
				SellerID:    "seller-1",
				SellerName:  "Store",
				Category:    "Test",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, nil
		},
		FindImagesByProductIDFunc: func(productID string) ([]entity.ProductImage, error) {
			return []entity.ProductImage{}, nil
		},
	}

	useCase := NewGetProductUseCase(mockRepo)
	result, err := useCase.Execute(ProductInputDTO{ID: "PROD-123"})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "PROD-123", result.ID)
	assert.Empty(t, result.Images)
}
