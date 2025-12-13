package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthResponse struct {
	Status    string    `json:"status" example:"healthy"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z"`
	Service   string    `json:"service" example:"product-api"`
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "product-api",
	})
}
