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
	assert.Equal(t, "The requested product was not found", response.Error)
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
	assert.Equal(t, "The provided product ID is invalid", response.Error)
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
	// For 500 errors, should return generic message
	assert.Equal(t, "An internal error occurred. Please try again later.", response.Error)
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
	// For 500 errors, should return generic message instead of exposing details
	assert.Equal(t, "An internal error occurred. Please try again later.", response.Error)
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

func TestErrorHandlerMiddleware_DoesNotLeakSensitiveInfo(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		shouldContain  string
		shouldNotContain []string
	}{
		{
			name:           "SQL error should not leak connection details",
			err:            fmt.Errorf("sql: connection failed - user:admin password:secret123 host:internal-db.local"),
			expectedStatus: http.StatusInternalServerError,
			shouldContain:  "An internal error occurred",
			shouldNotContain: []string{"admin", "secret123", "internal-db.local", "sql:"},
		},
		{
			name:           "panic should not leak stack trace",
			err:            fmt.Errorf("panic: runtime error: invalid memory address or nil pointer dereference"),
			expectedStatus: http.StatusInternalServerError,
			shouldContain:  "An internal error occurred",
			shouldNotContain: []string{"panic", "runtime error", "memory address"},
		},
		{
			name:           "database connection string should not be exposed",
			err:            fmt.Errorf("failed to connect: postgres://user:pass@db-server:5432/mydb"),
			expectedStatus: http.StatusInternalServerError,
			shouldContain:  "An internal error occurred",
			shouldNotContain: []string{"postgres://", "user:pass", "db-server"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupTestRouter()
			router.GET("/test", func(c *gin.Context) {
				_ = c.Error(tt.err)
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response errors.ErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Should contain safe message
			assert.Contains(t, response.Error, tt.shouldContain)

			// Should NOT contain sensitive information
			for _, sensitive := range tt.shouldNotContain {
				assert.NotContains(t, response.Error, sensitive,
					"Response should not leak sensitive information: %s", sensitive)
			}
		})
	}
}
