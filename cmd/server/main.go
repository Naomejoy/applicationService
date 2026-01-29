package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Naomejoy/app-service/internal/api"
	"github.com/Naomejoy/app-service/internal/db"
	"github.com/Naomejoy/app-service/internal/middleware"
	"github.com/Naomejoy/app-service/internal/repository"
	"github.com/Naomejoy/app-service/internal/service"
	"github.com/Naomejoy/app-service/pkg/config"

	"github.com/gin-gonic/gin"

	_ "github.com/Naomejoy/app-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Application Service API
// @version 1.0
// @description API for managing applications, statuses and file types
// @host localhost:8083
// @BasePath /api/v1
// @schemes http

// Define API key security
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key

// Apply security globally
// @security ApiKeyAuth
func main() {

	cfg := config.LoadConfig()
	log.Printf("Starting Application Service on port %s", cfg.Port)

	db.ConnectDB(cfg)

	appRepo := repository.NewApplicationRepository(db.DB)
	statusRepo := repository.NewApplicationStatusRepository(db.DB)
	fileRepo := repository.NewApplicationFileTypeRepository(db.DB)

	appService := service.NewApplicationService(appRepo)
	statusService := service.NewApplicationStatusService(statusRepo)
	fileService := service.NewApplicationFileTypeService(fileRepo)

	appHandler := api.NewApplicationHandler(appService)
	statusHandler := api.NewStatusHandler(statusService)
	fileHandler := api.NewFileTypeHandler(fileService)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "service is healthy"})
	})
	r.GET("/readyz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready", "message": "service is ready"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Application Service API",
			"docs":    "http://localhost:8083/swagger/index.html",
			"health":  "http://localhost:8083/healthz",
			"ready":   "http://localhost:8083/readyz",
		})
	})

	apiKey := cfg.APIKey

	log.Printf("API Key configured and is : %s", cfg.APIKey)

	api := r.Group("/api/v1")
	api.Use(middleware.APIKeyAuthMiddleware(apiKey))

	applications := api.Group("/applications")
	{
		applications.POST("", appHandler.CreateApplication)
		applications.GET("", appHandler.ListApplications)
		applications.GET("/:id", appHandler.GetApplication)
		applications.PUT("/:id", appHandler.UpdateApplication)
		applications.DELETE("/:id", appHandler.DeleteApplication)

		applications.POST("/:id/status", statusHandler.AddStatus)
		applications.GET("/:id/statuses", statusHandler.ListStatuses)

		applications.POST("/:id/file-types", fileHandler.AddFileType)
		applications.GET("/:id/file-types", fileHandler.ListFileTypes)
		applications.DELETE("/:id/file-types/:fileTypeId", fileHandler.DeleteFileType)
	}

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server started at http://localhost:%s", cfg.Port)
		log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}
