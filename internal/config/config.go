package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	AppPort    string
	GinMode    string
	APIVersion string
	APITimeout time.Duration
	LogLevel   string
	LogFormat  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables or defaults")
	}

	return &Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "8080"),
		GinMode:    getEnv("GIN_MODE", "debug"),
		APIVersion: getEnv("API_VERSION", "v1"),
		APITimeout: getEnvAsDuration("API_TIMEOUT", 30*time.Second),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		LogFormat:  getEnv("LOG_FORMAT", "text"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}
