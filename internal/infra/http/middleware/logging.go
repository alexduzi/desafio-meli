package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		requestID, _ := c.Get("request_id")

		logEvent := log.Info().
			Str("request_id", requestID.(string)).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Int("status", c.Writer.Status()).
			Dur("duration_ms", duration).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent())

		if len(c.Errors) > 0 {
			logEvent.Err(c.Errors.Last()).Msg("Request completed with errors")
		} else {
			if c.Writer.Status() >= 500 {
				logEvent.Msg("Request completed with server error")
			} else if c.Writer.Status() >= 400 {
				logEvent.Msg("Request completed with client error")
			} else {
				logEvent.Msg("Request completed successfully")
			}
		}
	}
}
