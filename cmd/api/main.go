package main

import (
	"log"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"

	_ "project/docs"
)

// @title Product API
// @version 1.0
// @description API for managing products
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := database.NewProductRepository(db)

	listProductUseCase := usecase.NewListProductUseCase(productRepo)
	getProductUseCase := usecase.NewGetProductUseCase(productRepo)

	productHandler := handler.NewProductHandler(listProductUseCase, getProductUseCase)

	router := httpInfra.SetupRouter(productHandler)

	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
