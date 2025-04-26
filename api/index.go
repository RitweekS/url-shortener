package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RitweekS/url-shortener.git/internal"
	"github.com/RitweekS/url-shortener.git/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Handler - Vercel serverless function handler
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set up Gin in serverless mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(internal.CORS())

	// Initialize database connection
	_ = godotenv.Load()
	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		fmt.Println("DBURL environment variable not set")
	}

	db, err := database.DBInit(dbURL)
	if err != nil {
		fmt.Printf("Database initialization error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	database.DB = db

	// Set up routes
	internal.InitializeRoutes(router)

	// Serve the request
	router.ServeHTTP(w, r)
}
