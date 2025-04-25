package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RitweekS/url-shortener.git/internal"
	"github.com/RitweekS/url-shortener.git/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		fmt.Println("DBURL environment variable not set")
	}

	db, err := database.DBInit(dbURL)
	if err != nil {
		fmt.Printf("Error initializing database: %v", err)
	}
	database.DB = db

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting DB instance: %v", err)
	}
	defer sqlDB.Close()

	// Initialize Gin router
	router = gin.New()
	router.Use(internal.CORS())
	internal.InitializeRoutes(router)
}

// Handler is the Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
