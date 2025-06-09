package handler

import (
	"net/http"

	"school_ride_backend/internal/model"
	"school_ride_backend/internal/service"

	"github.com/gin-gonic/gin"
)

// GetSchools godoc
// @Summary      List schools
// @Description  Get all schools
// @Tags         schools
// @Produce      json
// @Success      200  {array}  model.School
// @Router       /schools [get]
func CreateSchoolHandler(c *gin.Context) {
    var school model.School
    if err := c.ShouldBindJSON(&school); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.CreateSchool(c.Request.Context(), &school); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, school)
}
