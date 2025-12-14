package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"project/internal/dto"
	"project/internal/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockListProductUseCase struct {
	mock.Mock
}

func (m *MockListProductUseCase) Execute(ctx context.Context) ([]dto.ProductDTO, error) {
	args := m.Called(ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.ProductDTO), nil
}

type MockGetProductUseCase struct {
	mock.Mock
}

func (m *MockGetProductUseCase) Execute(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
	args := m.Called(ctx, input)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ProductDTO), nil
}

func setupTestRouter(handler *ProductHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			statusCode := errors.GetStatusCode(err)
			errorCode := errors.GetErrorCode(err)
			c.JSON(statusCode, errors.ErrorResponse{
				Error:     err.Error(),
				Code:      errorCode,
				Timestamp: time.Now(),
			})
		}
	})

	r.GET("/products", handler.ListProducts)
	r.GET("/products/:id", handler.GetProduct)

	return r
}

func TestProductHandler_ListProducts_Success(t *testing.T) {
	result := []dto.ProductDTO{
		{
			ID:       "PROD-1",
			Title:    "Product 1",
			Price:    100.0,
			Currency: "USD",
		},
		{
			ID:       "PROD-2",
			Title:    "Product 2",
			Price:    200.0,
			Currency: "USD",
		},
	}

	mockListUseCase := new(MockListProductUseCase)
	mockListUseCase.On("Execute", mock.Anything).Return(result, nil)

	handler := NewProductHandler(mockListUseCase, nil)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "data")

	data := response["data"].([]interface{})
	assert.Len(t, data, 2)
}

func TestProductHandler_ListProducts_EmptyList(t *testing.T) {
	mockListUseCase := new(MockListProductUseCase)
	mockListUseCase.On("Execute", mock.Anything).Return([]dto.ProductDTO{}, nil)

	handler := NewProductHandler(mockListUseCase, nil)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	data := response["data"].([]interface{})
	assert.Empty(t, data)
}

func TestProductHandler_ListProducts_DatabaseError(t *testing.T) {
	mockListUseCase := new(MockListProductUseCase)
	mockListUseCase.On("Execute", mock.Anything).Return(nil, fmt.Errorf("failed to list products: %w", errors.ErrDatabaseError))

	handler := NewProductHandler(mockListUseCase, nil)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DATABASE_ERROR", response.Code)
}

func TestProductHandler_GetProduct_Success(t *testing.T) {
	result := &dto.ProductDTO{
		ID:          "PROD-123",
		Title:       "iPhone 15",
		Description: "Latest iPhone",
		Price:       999.99,
		Currency:    "USD",
		Images: []dto.ProductImageDTO{
			{ID: 1, ImageURL: "http://example.com/img.jpg"},
		},
	}
	mockGetUseCase := new(MockGetProductUseCase)
	mockGetUseCase.On("Execute", mock.Anything, mock.Anything).Return(result, nil)

	handler := NewProductHandler(nil, mockGetUseCase)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/PROD-123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "data")

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "PROD-123", data["id"])
	assert.Equal(t, "iPhone 15", data["title"])
	assert.Equal(t, 999.99, data["price"])
}

func TestProductHandler_GetProduct_InvalidID(t *testing.T) {
	mockGetUseCase := new(MockGetProductUseCase)
	mockGetUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.ErrInvalidProductID)

	handler := NewProductHandler(nil, mockGetUseCase)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/   ", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_PRODUCT_ID", response.Code)
}

func TestProductHandler_GetProduct_NotFound(t *testing.T) {
	mockGetUseCase := new(MockGetProductUseCase)
	mockGetUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.ErrProductNotFound)

	handler := NewProductHandler(nil, mockGetUseCase)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/PROD-NONEXISTENT", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "PRODUCT_NOT_FOUND", response.Code)
	assert.Equal(t, "product not found", response.Error)
}

func TestProductHandler_GetProduct_DatabaseError(t *testing.T) {
	mockGetUseCase := new(MockGetProductUseCase)
	mockGetUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.ErrDatabaseError)

	handler := NewProductHandler(nil, mockGetUseCase)
	router := setupTestRouter(handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/PROD-123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DATABASE_ERROR", response.Code)
}
