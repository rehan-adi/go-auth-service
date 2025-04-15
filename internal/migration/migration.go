package main

import (
	"os"

	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

func RunMigrations() {

	// Initialize logger
	utils.InitLogger()

	// Load env variables
	if err := config.Init(); err != nil {
		utils.Log.Error("❌ Failed to load env variables", "error", err)
		os.Exit(1)
	}

	// Connect to database
	if err := database.ConnectDB(); err != nil {
		utils.Log.Error("❌ Failed to connect to the database", "error", err)
		os.Exit(1)
	}

	// Run migration
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		utils.Log.Error("❌ Migration failed", "error", err)
		os.Exit(1)
	}

	utils.Log.Info("✅ Database migration completed successfully")
}

func main() {
	RunMigrations()
}
