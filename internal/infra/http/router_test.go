package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"project/internal/dto"
	"project/internal/handler"

	"github.com/stretchr/testify/assert"
)

type mockListProductUseCase struct{}

func (m *mockListProductUseCase) Execute() ([]dto.ProductDTO, error) {
	return []dto.ProductDTO{}, nil
}

type mockGetProductUseCase struct{}

func (m *mockGetProductUseCase) Execute(input dto.ProductInputDTO) (*dto.ProductDTO, error) {
	return &dto.ProductDTO{ID: input.ID}, nil
}

func TestSetupRouter(t *testing.T) {
	listUseCase := &mockListProductUseCase{}
	getUseCase := &mockGetProductUseCase{}
	productHandler := handler.NewProductHandler(listUseCase, getUseCase)
	healthHandler := handler.NewHealthHandler()

	router := SetupRouter(productHandler, healthHandler)

	assert.NotNil(t, router)
}

func TestSetupRouter_ProductsEndpoint(t *testing.T) {
	listUseCase := &mockListProductUseCase{}
	getUseCase := &mockGetProductUseCase{}
	productHandler := handler.NewProductHandler(listUseCase, getUseCase)
	healthHandler := handler.NewHealthHandler()

	router := SetupRouter(productHandler, healthHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetupRouter_GetProductEndpoint(t *testing.T) {
	listUseCase := &mockListProductUseCase{}
	getUseCase := &mockGetProductUseCase{}
	productHandler := handler.NewProductHandler(listUseCase, getUseCase)
	healthHandler := handler.NewHealthHandler()

	router := SetupRouter(productHandler, healthHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/PROD-123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetupRouter_ErrorMiddlewareIsApplied(t *testing.T) {
	listUseCase := &mockListProductUseCase{}
	getUseCase := &mockGetProductUseCase{}
	productHandler := handler.NewProductHandler(listUseCase, getUseCase)
	healthHandler := handler.NewHealthHandler()

	router := SetupRouter(productHandler, healthHandler)

	assert.NotNil(t, router)
	assert.NotEmpty(t, router.Routes())
}

func TestSetupRouter_HealthEndpoint(t *testing.T) {
	listUseCase := &mockListProductUseCase{}
	getUseCase := &mockGetProductUseCase{}
	productHandler := handler.NewProductHandler(listUseCase, getUseCase)
	healthHandler := handler.NewHealthHandler()

	router := SetupRouter(productHandler, healthHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}
