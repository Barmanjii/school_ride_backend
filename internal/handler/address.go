package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAddressesHandler godoc
// @Summary		Get all addresses
// @Description	Retrieves a list of all addresses
// @Tags			addresses
// @Produce		json
// @Success		200
// @Failure		500
// @Router			/addresses [get]
func GetAddressesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}
