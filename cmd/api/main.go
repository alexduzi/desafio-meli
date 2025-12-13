package main

import (
	"fmt"
	"log"
	"project/internal/config"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"

	_ "project/docs"

	"github.com/gin-gonic/gin"
)

// @title Product API
// @version 1.0
// @description API for fetching product item
// @termsOfService http://swagger.io/terms/

// @contact.name Alex Duzi
// @contact.email duzihd@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := database.NewProductRepository(db)

	listProductUseCase := usecase.NewListProductUseCase(productRepo)
	getProductUseCase := usecase.NewGetProductUseCase(productRepo)

	productHandler := handler.NewProductHandler(listProductUseCase, getProductUseCase)
	healthHandler := handler.NewHealthHandler()

	router := httpInfra.SetupRouter(productHandler, healthHandler)

	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Starting server on %s (mode: %s, env: %s)", serverAddr, cfg.GinMode, cfg.AppEnv)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
