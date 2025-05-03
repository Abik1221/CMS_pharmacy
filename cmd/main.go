package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/your_project/config"
	"github.com/your_project/internal/routes"
)

func main() {
	// Load environment variables from .env file (optional but recommended)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found. Using system environment variables.")
	}

	// Connect to the database
	config.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Register all routes
	routes.RegisterRoutes(router)

	// Start server on specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default port
	}
	log.Printf("🚀 Server running at http://localhost:%s", port)
	router.Run(":" + port)
}
