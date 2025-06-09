package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz godoc
// @Summary      Health check endpoint
// @Description  Responds with a simple message to indicate the service is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string  "pong response"
// @Router       /healthz [get]
func Healthz(router *gin.Context) {
	router.JSON(http.StatusOK, gin.H{"message": "PONG"})
}
