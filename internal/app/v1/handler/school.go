package handler

import (
	"net/http"
	"school_ride_backend/internal/app/v1/model"
	"school_ride_backend/internal/app/v1/service"
	"school_ride_backend/internal/app/v1/utils"

	"github.com/gin-gonic/gin"
)

// GetSchoolsHandler godoc
//
//	@Description	Get all schools
//	@Tags			Schools
//	@Produce		json
//	@Success		200
//	@Router			/api/v1/schools/ [get]
func GetSchoolsHandler(c *gin.Context) {
	schools, err := service.ListAllSchools(c.Request.Context())
	if err != nil {
		utils.ResponseBody(c, 500, "Failed to retrieve schools: "+err.Error(), nil)
		return
	}
	if len(schools) == 0 {
		utils.ResponseBody(c, http.StatusOK, "No schools found", []model.School{})
		return
	}
	utils.ResponseBody(c, http.StatusOK, "Schools retrieved successfully", schools)
}

// CreateSchool godoc
//
//	@Description	Create a new school
//	@Tags			Schools
//	@Produce		json
//	@Param			school	body	model.School	true	"School object"
//	@Success		201		{array}	model.School
//	@Router			/api/v1/schools/ [post]
func CreateSchoolHandler(c *gin.Context) {
	var school model.School
	if err := c.ShouldBindJSON(&school); err != nil {
		utils.ResponseBody(c, 400, "Invalid request payload: "+err.Error(), nil)
		return
	}

	if err := service.CreateSchool(c.Request.Context(), &school); err != nil {
		utils.ResponseBody(c, http.StatusInternalServerError, "Failed to create school", nil)
		return
	}
	utils.ResponseBody(c, http.StatusCreated, "School created successfully", nil)
}
