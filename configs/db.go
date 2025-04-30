package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB initializes the MySQL database connection using GORM
func ConnectDB() {
	// Load from environment
	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASSWORD", "")
	host := GetEnv("DB_HOST", "127.0.0.1")
	port := GetEnv("DB_PORT", "3306")
	name := GetEnv("DB_NAME", "pharmacy")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	// GORM custom logger (optional)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("✅ Connected to the database successfully!")
}
