package main

import (
	"log"
	"os"
	"relynolli_oil_selection/core/api/brands"
	"relynolli_oil_selection/core/internal_services/database"
	"relynolli_oil_selection/core/internal_services/environment"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Create new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Relynolli Oil Selection",
	})
	// CORS configuration
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}
		
		return c.Next()
	})

	if (os.Getenv("ENV") == "development") {
		environment.SetupEnv()
	}

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	// Routes
	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8000"))
}

func setupRoutes(app *fiber.App) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Server is healthy",
		})
	})

	api := app.Group("/api")
	brands.SetupRoutes(api)
}
