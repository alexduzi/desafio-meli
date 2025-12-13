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
)

type MockListProductUseCase struct {
	ExecuteFunc func(ctx context.Context) ([]dto.ProductDTO, error)
}

func (m *MockListProductUseCase) Execute(ctx context.Context) ([]dto.ProductDTO, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx)
	}
	return []dto.ProductDTO{}, nil
}

type MockGetProductUseCase struct {
	ExecuteFunc func(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error)
}

func (m *MockGetProductUseCase) Execute(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx, input)
	}
	return nil, nil
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
	mockListUseCase := &MockListProductUseCase{
		ExecuteFunc: func(ctx context.Context) ([]dto.ProductDTO, error) {
			return []dto.ProductDTO{
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
			}, nil
		},
	}

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
	mockListUseCase := &MockListProductUseCase{
		ExecuteFunc: func(ctx context.Context) ([]dto.ProductDTO, error) {
			return []dto.ProductDTO{}, nil
		},
	}

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
	mockListUseCase := &MockListProductUseCase{
		ExecuteFunc: func(ctx context.Context) ([]dto.ProductDTO, error) {
			return nil, fmt.Errorf("failed to list products: %w", errors.ErrDatabaseError)
		},
	}

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
	mockGetUseCase := &MockGetProductUseCase{
		ExecuteFunc: func(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
			assert.Equal(t, "PROD-123", input.ID)
			return &dto.ProductDTO{
				ID:          "PROD-123",
				Title:       "iPhone 15",
				Description: "Latest iPhone",
				Price:       999.99,
				Currency:    "USD",
				Images: []dto.ProductImageDTO{
					{ID: 1, ImageURL: "http://example.com/img.jpg"},
				},
			}, nil
		},
	}

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
	mockGetUseCase := &MockGetProductUseCase{
		ExecuteFunc: func(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
			return nil, errors.ErrInvalidProductID
		},
	}

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
	mockGetUseCase := &MockGetProductUseCase{
		ExecuteFunc: func(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
			return nil, errors.ErrProductNotFound
		},
	}

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
	mockGetUseCase := &MockGetProductUseCase{
		ExecuteFunc: func(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
			return nil, fmt.Errorf("failed to get product: %w", errors.ErrDatabaseError)
		},
	}

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
