package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	// Routes
	routes.HealthRouter(server)

	log.Printf("🚀 Server running on port %s", config.Config.Port)
	server.Run(":" + config.Config.Port)

}
