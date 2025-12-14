package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/handler"
	"project/internal/infra/database"
	httpInfra "project/internal/infra/http"
	"project/internal/infra/logger"
	"project/internal/usecase"
	"syscall"
	"time"

	_ "project/docs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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

	logger.InitLogger(cfg.AppEnv)

	log.Info().
		Str("environment", cfg.AppEnv).
		Str("gin_mode", cfg.GinMode).
		Str("port", cfg.AppPort).
		Msg("Starting application")

	gin.SetMode(cfg.GinMode)

	db, err := database.InitDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	defer db.Close()

	log.Info().Msg("Database initialized successfully")

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
		log.Info().
			Str("address", serverAddr).
			Str("environment", cfg.AppEnv).
			Msg("Server starting")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	} else {
		log.Info().Msg("Server gracefully stopped")
	}
}
