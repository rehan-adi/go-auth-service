package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	JWT_SECRET  string
}

var AppConfig Config

func Init() error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	AppConfig = Config{
		Port:        getEnvOrPanic("PORT"),
		DB_PORT:     getEnvOrPanic("DB_PORT"),
		DB_HOST:     getEnvOrPanic("DB_HOST"),
		DB_USER:     getEnvOrPanic("DB_USER"),
		DB_NAME:     getEnvOrPanic("DB_NAME"),
		DB_PASSWORD: getEnvOrPanic("DB_PASSWORD"),
		JWT_SECRET:  getEnvOrPanic("JWT_SECRET"),
	}

	return nil
}

func getEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("❌ Missing required environment variable: %s", key))
	}
	return val
}
