package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/routes"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

func main() {

	// Initialize logger
	utils.InitLogger()

	// Load environment variables
	config.Init()
	utils.Log.Info("✅ Environment variables loaded successfully")

	// Initialize Gin server
	server := gin.Default()

	// Connect to database
	database.ConnectDB()

	// Middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := server.Group("/api/v1")

	// Routes
	routes.HealthRouter(api)
	routes.AuthRouter(api)
	routes.UserRouter(api)

	utils.Log.Infof("🚀 Server running on port %s", config.Port)

	if err := server.Run(":" + config.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

}
