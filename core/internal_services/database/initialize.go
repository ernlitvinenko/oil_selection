package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	// Get database connection details from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "fiber_db")
	dbPort := getEnv("DB_PORT", "5432")

	// Create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Connect to database
	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	log.Println("Connected Successfully to Database")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
