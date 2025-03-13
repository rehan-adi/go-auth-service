package config

import "os"

var Config = struct {
	Port        string
	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
}{
	Port:        getEnv("PORT", "3333"),
	DB_PORT:     getEnv("DB_PORT", "5432"),
	DB_HOST:     getEnv("DB_HOST", "localhost"),
	DB_USER:     getEnv("DB_USER", "user"),
	DB_NAME:     getEnv("DB_NAME", "database"),
	DB_PASSWORD: getEnv("DB_PASSWORD", "password"),
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
