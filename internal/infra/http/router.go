package http

import (
	"project/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	productHandler *handler.ProductHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/products", productHandler.ListProducts)
	}

	return r
}
