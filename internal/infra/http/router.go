package http

import (
	"project/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	productHandler *handler.ProductHandler,
	healthHandler *handler.HealthHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(ErrorHandlerMiddleware())

	r.GET("/health", healthHandler.HealthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.GET("/products", productHandler.ListProducts)
		api.GET("/products/:id", productHandler.GetProduct)
	}

	return r
}
