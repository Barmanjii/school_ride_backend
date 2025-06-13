// File: routes.go
// Package handler defines the HTTP handlers for the application.

package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// Register the routes for the application

	// NOTE: Vijay Barman 10th June - Not Creating the group for healthz endpoint
    router.GET("/healthz/ping", Healthz)


    v1 := router.Group("/api/v1")
    {
        // School routes
        school := v1.Group("/schools")
        {
			school.POST("/", CreateSchoolHandler)
            school.GET("/", GetSchoolsHandler)
        }
    }
}
