package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"pharmacy_cms/config"
	"pharmacy_cms/internal/routes"
)

func main() {
	// Load .env variables
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize DB
	db := config.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
