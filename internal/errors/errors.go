package errors

import (
	"errors"
	"net/http"
	"time"
)

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductID    = errors.New("invalid product id")
	ErrInvalidInput        = errors.New("invalid input")
	ErrDatabaseError       = errors.New("database error")
	ErrInternalServerError = errors.New("internal server error")
)

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

func NewAppError(err error, message string, statusCode int, code string) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		StatusCode: statusCode,
		Code:       code,
	}
}

type ErrorResponse struct {
	Error     string    `json:"error" example:"product not found"`
	Message   string    `json:"message,omitempty" example:"The requested product does not exist"`
	Code      string    `json:"code,omitempty" example:"PRODUCT_NOT_FOUND"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z"`
}

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

// GetUserFriendlyMessage returns a sanitized error message safe to expose to end users.
// For server errors (5xx), it returns a generic message to avoid leaking internal details.
// For client errors (4xx), it returns the custom message if available, or the error message.
func GetUserFriendlyMessage(err error, statusCode int) string {
	// For server errors (5xx), never expose internal details
	if statusCode >= 500 {
		var appErr *AppError
		if errors.As(err, &appErr) && appErr.Message != "" {
			return appErr.Message
		}
		return "An internal error occurred. Please try again later."
	}

	// For client errors (4xx), we can be more specific
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.Message != "" {
			return appErr.Message
		}
		// If AppError has no custom message, use the wrapped error
		if appErr.Err != nil {
			return appErr.Err.Error()
		}
	}

	// For known application errors, return their message
	switch {
	case errors.Is(err, ErrProductNotFound):
		return "The requested product was not found"
	case errors.Is(err, ErrInvalidProductID):
		return "The provided product ID is invalid"
	case errors.Is(err, ErrInvalidInput):
		return "The request contains invalid input"
	case errors.Is(err, ErrDatabaseError):
		return "An error occurred while accessing the database"
	default:
		// For unknown client errors, return a generic message
		return err.Error()
	}
}
