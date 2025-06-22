package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseBody sends a generic JSON response with custom message, data, and status code.
func ResponseBody(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
		"sucess":  statusCode >= http.StatusOK && statusCode < http.StatusBadRequest,
	})
}