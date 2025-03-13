package database

import (
	"fmt"

	"github.com/rehan-adi/go-auth-service/config"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		utils.Log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	utils.Log.Info("✅ Database connected successfully")

	DB = db

	// Auto Migrate the Model
	if !DB.Migrator().HasTable(&models.User{}) {
		err = DB.AutoMigrate(&models.User{})
		if err != nil {
			utils.Log.Fatalf("❌ Failed to migrate the database: %v", err)
		}
		utils.Log.Info("✅ Database migrated successfully")
	}

}
