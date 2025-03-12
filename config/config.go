package config

import "os"

var Config = struct {
	Port string
}{
	Port: getEnv("PORT", "3333"),
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
