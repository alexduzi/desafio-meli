package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/usecase"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router.Handler(),
	}

	go func() {
		log.Printf("starting server on %s (mode: %s, env: %s)\n", serverAddr, cfg.GinMode, cfg.AppEnv)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v\n", err)
	}
	fmt.Println("server exiting")
}
