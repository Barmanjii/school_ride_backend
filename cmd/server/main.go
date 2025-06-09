package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"school_ride_backend/internal/config"
	"school_ride_backend/internal/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, continuing...")
	}

	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()
	handler.RegisterRoutes(r)

	port := os.Getenv("PORT")
	// Default to port 8080 if not set
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
