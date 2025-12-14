package logger

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(env string) {
	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}

func FromContext(c *gin.Context) *zerolog.Logger {
	logger := log.With().Logger()

	if requestID, exists := c.Get("request_id"); exists {
		logger = logger.With().Str("request_id", requestID.(string)).Logger()
	}

	logger = logger.With().
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("client_ip", c.ClientIP()).
		Logger()

	return &logger
}

func GetLogger() *zerolog.Logger {
	return &log.Logger
}

type LogWriter struct {
	Logger *zerolog.Logger
}

func (w LogWriter) Write(p []byte) (n int, err error) {
	w.Logger.Info().Msg(string(p))
	return len(p), nil
}

func NewLogWriter(logger *zerolog.Logger) io.Writer {
	return LogWriter{Logger: logger}
}
