package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"

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

func TestIntegration_ListProducts(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
}

func TestIntegration_GetProduct_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/invalid-id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "PRODUCT_NOT_FOUND")
}

func TestIntegration_GetProduct_EmptyID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/   ", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
