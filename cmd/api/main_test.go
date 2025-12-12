package main

import (
	"net/http"
	"net/http/httptest"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	db, err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	productRepo := database.NewProductRepository(db)
	listProductUseCase := usecase.NewListProductUseCase(productRepo)
	getProductUseCase := usecase.NewGetProductUseCase(productRepo)

	productHandler := handler.NewProductHandler(listProductUseCase, getProductUseCase)

	return httpInfra.SetupRouter(productHandler)
}

func TestListProducts(t *testing.T) {
	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
}

func TestGetProduct(t *testing.T) {
	router := setupTestRouter(t)

	// Test with invalid ID to check error handling
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/invalid-id", nil)
	router.ServeHTTP(w, req)

	// Should return 404 or 500 depending on the error
	assert.NotEqual(t, http.StatusOK, w.Code)
}
