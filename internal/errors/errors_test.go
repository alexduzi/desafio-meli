package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusCode(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
	}{
		{
			name:           "Product not found returns 404",
			err:            ErrProductNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid product ID returns 400",
			err:            ErrInvalidProductID,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid input returns 400",
			err:            ErrInvalidInput,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Database error returns 500",
			err:            ErrDatabaseError,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Unknown error returns 500",
			err:            errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "AppError with custom status code",
			err:            NewAppError(nil, "custom error", http.StatusConflict, "CONFLICT"),
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := GetStatusCode(tt.err)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}

func TestGetErrorCode(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode string
	}{
		{
			name:         "Product not found",
			err:          ErrProductNotFound,
			expectedCode: "PRODUCT_NOT_FOUND",
		},
		{
			name:         "Invalid product ID",
			err:          ErrInvalidProductID,
			expectedCode: "INVALID_PRODUCT_ID",
		},
		{
			name:         "Invalid input",
			err:          ErrInvalidInput,
			expectedCode: "INVALID_INPUT",
		},
		{
			name:         "Database error",
			err:          ErrDatabaseError,
			expectedCode: "DATABASE_ERROR",
		},
		{
			name:         "Unknown error",
			err:          errors.New("unknown"),
			expectedCode: "INTERNAL_ERROR",
		},
		{
			name:         "AppError with custom code",
			err:          NewAppError(nil, "custom", http.StatusBadRequest, "CUSTOM_CODE"),
			expectedCode: "CUSTOM_CODE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := GetErrorCode(tt.err)
			assert.Equal(t, tt.expectedCode, code)
		})
	}
}

func TestAppError(t *testing.T) {
	t.Run("Error method returns underlying error message", func(t *testing.T) {
		underlyingErr := errors.New("database connection failed")
		appErr := NewAppError(underlyingErr, "custom message", http.StatusInternalServerError, "DB_ERROR")

		assert.Equal(t, "database connection failed", appErr.Error())
	})

	t.Run("Error method returns message when no underlying error", func(t *testing.T) {
		appErr := NewAppError(nil, "custom message", http.StatusBadRequest, "VALIDATION_ERROR")

		assert.Equal(t, "custom message", appErr.Error())
	})

	t.Run("Unwrap returns underlying error", func(t *testing.T) {
		underlyingErr := errors.New("original error")
		appErr := NewAppError(underlyingErr, "wrapped", http.StatusInternalServerError, "ERROR")

		assert.Equal(t, underlyingErr, appErr.Unwrap())
	})

	t.Run("AppError fields are set correctly", func(t *testing.T) {
		underlyingErr := errors.New("test error")
		appErr := NewAppError(underlyingErr, "test message", http.StatusNotFound, "TEST_CODE")

		assert.Equal(t, underlyingErr, appErr.Err)
		assert.Equal(t, "test message", appErr.Message)
		assert.Equal(t, http.StatusNotFound, appErr.StatusCode)
		assert.Equal(t, "TEST_CODE", appErr.Code)
	})
}

func TestWrappedErrors(t *testing.T) {
	t.Run("GetStatusCode works with wrapped errors", func(t *testing.T) {
		wrappedErr := errors.Join(ErrProductNotFound, errors.New("additional context"))
		status := GetStatusCode(wrappedErr)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("GetErrorCode works with wrapped errors", func(t *testing.T) {
		wrappedErr := errors.Join(ErrInvalidProductID, errors.New("validation failed"))
		code := GetErrorCode(wrappedErr)
		assert.Equal(t, "INVALID_PRODUCT_ID", code)
	})
}
