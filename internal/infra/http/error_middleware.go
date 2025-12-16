package http

import (
	"time"

	"project/internal/errors"
	"project/internal/infra/logger"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			log := logger.FromContext(c)

			statusCode := errors.GetStatusCode(err)
			errorCode := errors.GetErrorCode(err)

			// Log the full error details internally (including sensitive information)
			if statusCode >= 500 {
				log.Error().
					Err(err).
					Str("error_code", errorCode).
					Int("status_code", statusCode).
					Str("error_details", err.Error()).
					Msg("Server error occurred")
			} else if statusCode >= 400 {
				log.Warn().
					Err(err).
					Str("error_code", errorCode).
					Int("status_code", statusCode).
					Msg("Client error occurred")
			}

			// Return sanitized error message to the user
			userMessage := errors.GetUserFriendlyMessage(err, statusCode)
			errorResponse := errors.ErrorResponse{
				Error:     userMessage,
				Code:      errorCode,
				Timestamp: time.Now(),
			}

			c.JSON(statusCode, errorResponse)
		}
	}
}
