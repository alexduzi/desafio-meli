package usecase

import (
	"fmt"
	"testing"
	"time"

	"project/internal/entity"
	"project/internal/errors"
	"project/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListProductUseCaseTestSuite struct {
	suite.Suite
	repositoryMock *repository.MockProductRepository
}

func (suite *ListProductUseCaseTestSuite) BeforeTest(suiteName, testName string) {
	suite.repositoryMock = new(repository.MockProductRepository)
}

func (suite *ListProductUseCaseTestSuite) TestListProductUseCase_Execute_Success() {
	now := time.Now()
	products := []entity.Product{
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
			Thumbnail:   "https://example.com/thumbnails/prod1.jpg",
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
			Thumbnail:   "https://example.com/thumbnails/prod2.jpg",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	suite.repositoryMock.On("ListProducts").Return(products, nil)

	useCase := NewListProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute()

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), "PROD-1", result[0].ID)
	assert.Equal(suite.T(), "Product 1", result[0].Title)
	assert.Equal(suite.T(), 100.0, result[0].Price)
	assert.Equal(suite.T(), "https://example.com/thumbnails/prod1.jpg", result[0].Thumbnail)
	assert.Equal(suite.T(), "PROD-2", result[1].ID)
	assert.Equal(suite.T(), "https://example.com/thumbnails/prod2.jpg", result[1].Thumbnail)
}

func (suite *ListProductUseCaseTestSuite) TestListProductUseCase_Execute_EmptyList() {
	products := []entity.Product{}

	suite.repositoryMock.On("ListProducts").Return(products, nil)

	useCase := NewListProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute()

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Empty(suite.T(), result)
}

func (suite *ListProductUseCaseTestSuite) TestListProductUseCase_Execute_DatabaseError() {
	suite.repositoryMock.On("ListProducts").Return(nil, fmt.Errorf("%w: connection timeout", errors.ErrDatabaseError))

	useCase := NewListProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute()

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.ErrorIs(suite.T(), err, errors.ErrDatabaseError)
}

func (suite *ListProductUseCaseTestSuite) TestListProductUseCase_Execute_MapsAllFields() {
	now := time.Now()
	products := []entity.Product{
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
			Thumbnail:   "https://cdn.example.com/thumb-test.jpg",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	suite.repositoryMock.On("ListProducts").Return(products, nil)

	useCase := NewListProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute()

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 1)

	product := result[0]
	assert.Equal(suite.T(), "PROD-TEST", product.ID)
	assert.Equal(suite.T(), "Test Title", product.Title)
	assert.Equal(suite.T(), "Test Description", product.Description)
	assert.Equal(suite.T(), 99.99, product.Price)
	assert.Equal(suite.T(), "EUR", product.Currency)
	assert.Equal(suite.T(), "refurbished", product.Condition)
	assert.Equal(suite.T(), 3, product.Stock)
	assert.Equal(suite.T(), "seller-test", product.SellerID)
	assert.Equal(suite.T(), "Test Seller", product.SellerName)
	assert.Equal(suite.T(), "Test Category", product.Category)
	assert.Equal(suite.T(), "https://cdn.example.com/thumb-test.jpg", product.Thumbnail)
	assert.Equal(suite.T(), now, product.CreatedAt)
	assert.Equal(suite.T(), now, product.UpdatedAt)
}

func (suite *ListProductUseCaseTestSuite) TestListProductUseCase_Execute_DoesNotIncludeImages() {
	now := time.Now()
	products := []entity.Product{
		{
			ID:        "PROD-1",
			Title:     "Product with Thumbnail",
			Price:     50.0,
			Currency:  "USD",
			Condition: "new",
			Stock:     5,
			SellerID:  "seller-1",
			Thumbnail: "https://example.com/thumb.jpg",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	suite.repositoryMock.On("ListProducts").Return(products, nil)

	useCase := NewListProductUseCase(suite.repositoryMock)
	result, err := useCase.Execute()

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 1)
	assert.Equal(suite.T(), "https://example.com/thumb.jpg", result[0].Thumbnail)
	assert.Nil(suite.T(), result[0].Images)
}

func TestListProductUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ListProductUseCaseTestSuite))
}
