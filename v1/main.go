package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"school_ride_backend/v1/config"
	_ "school_ride_backend/v1/docs"
	"school_ride_backend/v1/handler"

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

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, continuing...")
	}

	config.InitDB()
	defer config.DB.Close()

	// NOTE (Vijay barman, 19th Jun 2025): Following the docs - https://github.com/gin-gonic/gin/blob/master/docs/doc.md#using-middleware

	// Creates a router without any middleware by default
	router := gin.New()

	// NOTE (Vijay Barman, 23rd June 2025): This is simply like cross-origin resource sharing (CORS) in the browser.
	//   router.SetTrustedProxies([]string{"192.168.1.2"}) // Set trusted proxies if needed, otherwise remove this line

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	handler.RegisterRoutes(router)

	port := os.Getenv("PORT")

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Default to port 8080 if not set
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
