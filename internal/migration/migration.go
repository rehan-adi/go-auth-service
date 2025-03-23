package main

import (
	"fmt"
	"log"

	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

func RunMigrations() {

	// Initialize logger
	utils.InitLogger()

	// Load env variables
	config.Init()

	// Connect to database
	database.ConnectDB()

	// Run migration
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	fmt.Println("✅ Database migration completed successfully")
}

func main() {
	RunMigrations()
}
