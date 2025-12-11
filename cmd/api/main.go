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
		panic(err)
	}

	productRepo := database.NewProductRepository(db)

	listProductUseCase := usecase.NewListProductUseCase(productRepo)
	getProductUseCase := usecase.NewGetProductUseCase(productRepo)

	productHandler := handler.NewProductHandler(listProductUseCase, getProductUseCase)

	router := httpInfra.SetupRouter(productHandler)

	if err := router.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
