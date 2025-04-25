package main

import (
	"fmt"
	"os"

	"github.com/RitweekS/url-shortener.git/internal"
	"github.com/RitweekS/url-shortener.git/internal/database"
	"github.com/RitweekS/url-shortener.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		fmt.Println("DBURL environment variable not set")
	}
	db, err := database.DBInit(dbURL)
	database.DB = db
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting DB instance: %v", err)
	}
	autoMigErr := db.AutoMigrate(&models.Shortener{})
	if autoMigErr != nil {
		fmt.Printf("Failed to migrate database schema: %v", err)
	}

	defer sqlDB.Close()
	router := gin.Default()
	router.Use(internal.CORS())
	internal.InitializeRoutes(router)

	router.Run(":8081")
}
