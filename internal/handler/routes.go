// File: routes.go
// Package handler defines the HTTP handlers for the application.

package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// Register the routes for the application
    router.GET("/healthz/ping", Healthz)
}
