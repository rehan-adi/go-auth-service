package database

import (
	"fmt"
	"log"
	"os"

	"github.com/rehan-adi/go-auth-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	fmt.Println("✅ Database connected successfully")
	DB = db

	// Auto Migrate the Model
	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("❌ Failed to migrate the database:", err)
	}

	fmt.Println("✅ Database migrated successfully")

}
