package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetUserFriendlyMessage_ServerErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		statusCode int
		want       string
	}{
		{
			name:       "generic 500 error should return generic message",
			err:        errors.New("sql: database connection failed with credentials user:password@localhost"),
			statusCode: http.StatusInternalServerError,
			want:       "An internal error occurred. Please try again later.",
		},
		{
			name:       "database error should return generic message for 500",
			err:        ErrDatabaseError,
			statusCode: http.StatusInternalServerError,
			want:       "An internal error occurred. Please try again later.",
		},
		{
			name: "AppError with custom message should use it for 500",
			err: &AppError{
				Err:        errors.New("connection pool exhausted"),
				Message:    "Service temporarily unavailable",
				StatusCode: http.StatusServiceUnavailable,
				Code:       "SERVICE_UNAVAILABLE",
			},
			statusCode: http.StatusServiceUnavailable,
			want:       "Service temporarily unavailable",
		},
		{
			name: "AppError without custom message should return generic for 500",
			err: &AppError{
				Err:        errors.New("panic: runtime error"),
				StatusCode: http.StatusInternalServerError,
				Code:       "INTERNAL_ERROR",
			},
			statusCode: http.StatusInternalServerError,
			want:       "An internal error occurred. Please try again later.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUserFriendlyMessage(tt.err, tt.statusCode)
			if got != tt.want {
				t.Errorf("GetUserFriendlyMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserFriendlyMessage_ClientErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		statusCode int
		want       string
	}{
		{
			name:       "product not found should return friendly message",
			err:        ErrProductNotFound,
			statusCode: http.StatusNotFound,
			want:       "The requested product was not found",
		},
		{
			name:       "invalid product ID should return friendly message",
			err:        ErrInvalidProductID,
			statusCode: http.StatusBadRequest,
			want:       "The provided product ID is invalid",
		},
		{
			name:       "invalid input should return friendly message",
			err:        ErrInvalidInput,
			statusCode: http.StatusBadRequest,
			want:       "The request contains invalid input",
		},
		{
			name: "AppError with custom message should use it for 4xx",
			err: &AppError{
				Err:        ErrInvalidInput,
				Message:    "Name field is required",
				StatusCode: http.StatusBadRequest,
				Code:       "INVALID_INPUT",
			},
			statusCode: http.StatusBadRequest,
			want:       "Name field is required",
		},
		{
			name: "AppError with wrapped error should show wrapped error for 4xx",
			err: &AppError{
				Err:        ErrProductNotFound,
				StatusCode: http.StatusNotFound,
				Code:       "PRODUCT_NOT_FOUND",
			},
			statusCode: http.StatusNotFound,
			want:       "product not found",
		},
		{
			name:       "unknown client error should return error message",
			err:        errors.New("some validation error"),
			statusCode: http.StatusBadRequest,
			want:       "some validation error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUserFriendlyMessage(tt.err, tt.statusCode)
			if got != tt.want {
				t.Errorf("GetUserFriendlyMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserFriendlyMessage_DoesNotLeakSensitiveInfo(t *testing.T) {
	sensitiveErrors := []error{
		errors.New("sql: connection failed - username: admin, password: secret123"),
		errors.New("panic: runtime error: invalid memory address or nil pointer dereference"),
		errors.New("database connection string: postgres://user:pass@internal-db:5432/mydb"),
		errors.New("stack trace: goroutine 1 [running]: main.main()"),
	}

	for _, err := range sensitiveErrors {
		got := GetUserFriendlyMessage(err, http.StatusInternalServerError)

		// Should not contain sensitive information
		if got == err.Error() {
			t.Errorf("GetUserFriendlyMessage() leaked sensitive error: %v", err.Error())
		}

		// Should return generic message
		if got != "An internal error occurred. Please try again later." {
			t.Errorf("GetUserFriendlyMessage() = %v, want generic message", got)
		}
	}
}
