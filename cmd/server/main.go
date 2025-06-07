package main

import (
	"school_ride_backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default() // Create a new Gin router instance
	router.Use(gin.Logger()) // Use the logger middleware for logging requests
    handler.RegisterRoutes(router) // Register the routes defined in the handler package
    router.Run(":8000") // Start the server on port 8000
}
