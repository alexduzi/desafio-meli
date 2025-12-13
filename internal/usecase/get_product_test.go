package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"project/internal/dto"
	"project/internal/entity"
	"project/internal/errors"
	"project/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetProductUseCaseTestSuite struct {
	suite.Suite
	repositoryMock *repository.MockProductRepository
}

func (suite *GetProductUseCaseTestSuite) BeforeTest(suiteName, testName string) {
	suite.repositoryMock = new(repository.MockProductRepository)
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_Success() {
	product := &entity.Product{
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
	}
	images := []entity.ProductImage{
		{ID: 1, ProductID: "PROD-123", ImageURL: "http://example.com/image1.jpg", DisplayOrder: 0},
		{ID: 2, ProductID: "PROD-123", ImageURL: "http://example.com/image2.jpg", DisplayOrder: 1},
	}

	suite.repositoryMock.On("GetProduct", mock.Anything, mock.Anything).Return(product, nil)

	suite.repositoryMock.On("FindImagesByProductID", mock.Anything, mock.Anything).Return(images, nil)

	useCase := NewGetProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: "PROD-123"})

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "PROD-123", result.ID)
	assert.Equal(suite.T(), "iPhone 15", result.Title)
	assert.Equal(suite.T(), 999.99, result.Price)
	assert.Len(suite.T(), result.Images, 2)
	assert.Equal(suite.T(), "http://example.com/image1.jpg", result.Images[0].ImageURL)
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_EmptyID() {
	useCase := NewGetProductUseCase(suite.repositoryMock)

	tests := []struct {
		name string
		id   string
	}{
		{"Empty string", ""},
		{"Only spaces", "   "},
		{"Only tabs", "\t\t"},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: tt.id})

			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Equal(t, errors.ErrInvalidProductID, err)
		})
	}
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_ProductNotFound() {
	suite.repositoryMock.On("GetProduct", mock.Anything, mock.Anything).Return(nil, errors.ErrProductNotFound)

	useCase := NewGetProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: "PROD-999"})

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.ErrorIs(suite.T(), err, errors.ErrProductNotFound)
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_DatabaseError() {
	suite.repositoryMock.On("GetProduct", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("%w: connection failed", errors.ErrDatabaseError))

	useCase := NewGetProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: "PROD-123"})

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.ErrorIs(suite.T(), err, errors.ErrDatabaseError)
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_ImagesError() {
	product := &entity.Product{
		ID:    "PROD-123",
		Title: "Test Product",
	}

	suite.repositoryMock.On("GetProduct", mock.Anything, mock.Anything).Return(product, nil)

	suite.repositoryMock.On("FindImagesByProductID", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("%w: failed to fetch images", errors.ErrDatabaseError))

	useCase := NewGetProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: "PROD-123"})

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.ErrorIs(suite.T(), err, errors.ErrDatabaseError)
}

func (suite *GetProductUseCaseTestSuite) TestGetProductUseCase_Execute_NoImages() {
	product := &entity.Product{
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
	}

	images := []entity.ProductImage{}

	suite.repositoryMock.On("GetProduct", mock.Anything, mock.Anything).Return(product, nil)

	suite.repositoryMock.On("FindImagesByProductID", mock.Anything, mock.Anything).Return(images, nil)

	useCase := NewGetProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute(context.Background(), dto.ProductInputDTO{ID: "PROD-123"})

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "PROD-123", result.ID)
	assert.Empty(suite.T(), result.Images)
}

func TestGetProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetProductUseCaseTestSuite))
}
