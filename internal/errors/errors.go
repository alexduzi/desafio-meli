package errors

import (
	"errors"
	"net/http"
	"time"
)

// Domain errors
var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductID    = errors.New("invalid product id")
	ErrInvalidInput        = errors.New("invalid input")
	ErrDatabaseError       = errors.New("database error")
	ErrInternalServerError = errors.New("internal server error")
)

// AppError represents an application error with HTTP status code
type AppError struct {
	Err        error
	Message    string
	StatusCode int
	Code       string
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError
func NewAppError(err error, message string, statusCode int, code string) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		StatusCode: statusCode,
		Code:       code,
	}
}

// ErrorResponse represents the JSON error response structure
type ErrorResponse struct {
	Error     string    `json:"error"`
	Message   string    `json:"message,omitempty"`
	Code      string    `json:"code,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// GetStatusCode returns the appropriate HTTP status code for an error
func GetStatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.StatusCode
	}

	switch {
	case errors.Is(err, ErrProductNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrInvalidProductID), errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, ErrDatabaseError):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// GetErrorCode returns a string code for the error type
func GetErrorCode(err error) string {
	var appErr *AppError
	if errors.As(err, &appErr) && appErr.Code != "" {
		return appErr.Code
	}

	switch {
	case errors.Is(err, ErrProductNotFound):
		return "PRODUCT_NOT_FOUND"
	case errors.Is(err, ErrInvalidProductID):
		return "INVALID_PRODUCT_ID"
	case errors.Is(err, ErrInvalidInput):
		return "INVALID_INPUT"
	case errors.Is(err, ErrDatabaseError):
		return "DATABASE_ERROR"
	default:
		return "INTERNAL_ERROR"
	}
}
