package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(router *gin.Context) {
	router.JSON(http.StatusOK, gin.H{"message": "PONG"})
}
