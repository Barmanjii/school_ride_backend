package handler

import (
	"net/http"
	"school_ride_backend/internal/app/v1/model"
	"school_ride_backend/internal/app/v1/service"
	"school_ride_backend/internal/app/v1/utils"

	"github.com/gin-gonic/gin"
)

// GetParentsHandler godoc
// @Summary		Get all parents
// @Description	Retrieves a list of all parents
// @Tags			parents
// @Produce		json
// @Success		200
// @Failure		500
// @Router			/api/v1/parents/ [get]
func GetParentsHandler(c *gin.Context) {
	parents, err := service.ListParents(c.Request.Context())
	if err != nil {
		utils.ResponseBody(c, 500, "Failed to retrieve parents: "+err.Error(), nil)
		return
	}

	if len(parents) == 0 {
		utils.ResponseBody(c, http.StatusOK, "No parents found", []model.ParentInDatabase{})
		return
	}
	utils.ResponseBody(c, 200, "Parents retrieved successfully", parents)
}

// CreateParentHandler godoc
// @Summary		Create a new parent
// @Description	Creates a new parent with the provided details
// @Tags			parents
// @Accept			json
// @Produce		json
// @Param			parent	body	model.Parent	true	"Parent details"
// @Success		201
// @Failure		400
// @Failure		500
// @Router			/api/v1/parents/ [post]
func CreateParentHandler(c *gin.Context) {
	var parent model.Parent

	// Validate input
	if err := c.ShouldBindJSON(&parent); err != nil {
		utils.ResponseBody(c, 400, "Invalid request payload: "+err.Error(), nil)
		return
	}

	// Insert into DB
	if err := service.CreateParent(c.Request.Context(), &parent); err != nil {
		utils.ResponseBody(c, 500, "Failed to create parent: "+err.Error(), nil)
		return
	}
	utils.ResponseBody(c, 201, "Parent created successfully", nil)
}
