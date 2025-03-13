package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

var (
	Port        string
	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
)

func Init() {

	err := godotenv.Load()

	if err != nil {
		utils.Log.Fatalf("Error loading .env file: %v", err)
	}

	Port = os.Getenv("PORT")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	utils.Log.Info("✅ Environment variables loaded successfully")
}
