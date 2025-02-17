package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvWithFallback(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// setup env variables using godotenv
func SetupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}