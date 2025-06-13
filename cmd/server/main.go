package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"school_ride_backend/internal/config"
	"school_ride_backend/internal/handler"

	_ "school_ride_backend/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	School Ride Backend Server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, continuing...")
	}

	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()
	handler.RegisterRoutes(r)

	port := os.Getenv("PORT")

	 // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Default to port 8080 if not set
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
