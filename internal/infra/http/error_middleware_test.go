package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"project/internal/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ErrorHandlerMiddleware())
	return r
}

func TestErrorHandlerMiddleware_NoError(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])
}

func TestErrorHandlerMiddleware_ProductNotFound(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(errors.ErrProductNotFound)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "product not found", response.Error)
	assert.Equal(t, "PRODUCT_NOT_FOUND", response.Code)
	assert.False(t, response.Timestamp.IsZero())
}

func TestErrorHandlerMiddleware_InvalidProductID(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(errors.ErrInvalidProductID)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "invalid product id", response.Error)
	assert.Equal(t, "INVALID_PRODUCT_ID", response.Code)
}

func TestErrorHandlerMiddleware_DatabaseError(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(fmt.Errorf("%w: connection failed", errors.ErrDatabaseError))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DATABASE_ERROR", response.Code)
	assert.Contains(t, response.Error, "database error")
}

func TestErrorHandlerMiddleware_UnknownError(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(fmt.Errorf("something went wrong"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "something went wrong", response.Error)
	assert.Equal(t, "INTERNAL_ERROR", response.Code)
}

func TestErrorHandlerMiddleware_CustomAppError(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		customErr := errors.NewAppError(
			nil,
			"custom error message",
			http.StatusConflict,
			"CUSTOM_CODE",
		)
		_ = c.Error(customErr)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "custom error message", response.Error)
	assert.Equal(t, "CUSTOM_CODE", response.Code)
}

func TestErrorHandlerMiddleware_MultipleErrors(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(errors.ErrInvalidProductID)
		_ = c.Error(errors.ErrProductNotFound)
		_ = c.Error(errors.ErrDatabaseError)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DATABASE_ERROR", response.Code)
}

func TestErrorHandlerMiddleware_ResponseFormat(t *testing.T) {
	router := setupTestRouter()
	router.GET("/test", func(c *gin.Context) {
		_ = c.Error(errors.ErrProductNotFound)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	var response errors.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.Error)
	assert.NotEmpty(t, response.Code)
	assert.False(t, response.Timestamp.IsZero())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
