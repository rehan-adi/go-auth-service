package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/middlewares"
	"github.com/rehan-adi/go-auth-service/internal/routes"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

func main() {

	// Initialize logger
	utils.InitLogger()

	// Load environment variables
	if err := config.Init(); err != nil {
		utils.Log.Error("❌ Failed to load env", "error", err)
		os.Exit(1)
	}

	utils.Log.Info("✅ Environment variables loaded successfully")

	// Initialize Gin server
	server := gin.Default()

	// Connect to database
	if err := database.ConnectDB(); err != nil {
		utils.Log.Error("❌ Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.Use(middlewares.RateLimiterMiddleware())

	api := server.Group("/api/v1")

	// Routes
	routes.HealthRouter(api)
	routes.AuthRouter(api)
	routes.UserRouter(api)

	utils.Log.Info("🚀 Server is running", "port", config.AppConfig.Port)

	if err := server.Run(":" + config.AppConfig.Port); err != nil {
		utils.Log.Error("Failed to start server", "error", err)
	}

}
