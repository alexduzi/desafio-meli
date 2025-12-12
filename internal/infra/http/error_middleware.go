package http

import (
	"log"
	"time"

	"project/internal/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			log.Printf("Error occurred: %v", err)

			statusCode := errors.GetStatusCode(err)
			errorCode := errors.GetErrorCode(err)

			errorResponse := errors.ErrorResponse{
				Error:     err.Error(),
				Code:      errorCode,
				Timestamp: time.Now(),
			}

			c.JSON(statusCode, errorResponse)
		}
	}
}
