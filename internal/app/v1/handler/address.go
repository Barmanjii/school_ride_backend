package handler

import (
	"school_ride_backend/internal/app/v1/model"
	"school_ride_backend/internal/app/v1/service"
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

// CreateAddressHandler godoc
//
//	@Summary		Create a new address
//	@Description	Creates a new address with the provided details
//	@Tags			addresses
//	@Accept			json
//	@Produce		json
//	@Param			address	body	model.Address	true	"Address details"
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/addresses/ [post]
func CreateAddressHandler(c *gin.Context) {
	var address model.Address

	// Validate input
	if err := c.ShouldBindJSON(&address); err != nil {
		utils.ResponseBody(c, 400, "Invalid request payload", nil)
		return
	}

	// Insert into DB
	if err := service.CreateAddress(c.Request.Context(), &address); err != nil {
		utils.ResponseBody(c, 500, "Failed to create address", nil)
		return
	}

	utils.ResponseBody(c, 201, "Address created successfully", nil)
}
