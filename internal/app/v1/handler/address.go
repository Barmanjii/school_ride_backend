package handler

import (
	"school_ride_backend/internal/app/v1/utils"

	"github.com/gin-gonic/gin"
)

// GetAddressesHandler godoc
//
//	@Summary		Get all addresses
//	@Description	Retrieves a list of all addresses
//	@Tags			addresses
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/addresses/ [get]
func GetAddressesHandler(c *gin.Context) {
	utils.ResponseBody(c, 200, "GetAddressesHandler not implemented yet", nil)
}
