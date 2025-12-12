package main

import (
	"log"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"
)

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
