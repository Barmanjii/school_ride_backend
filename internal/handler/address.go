package handler

import (
	"net/http"

	"school_ride_backend/internal/model"
	"school_ride_backend/internal/service"
	"school_ride_backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetAddressesHandler(c *gin.Context) {
	addresses, err := service.ListAddresses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(addresses) == 0 {
		utils.ResponseBody(c, http.StatusOK, "No addresses found", []model.Address{})
		return
	}
	utils.ResponseBody(c, http.StatusOK, "Addresses retrieved successfully", addresses)
}
