package handler

import (
	"net/http"
	"school_ride_backend/internal/app/v1/model"
	"school_ride_backend/internal/app/v1/service"
	"school_ride_backend/internal/app/v1/utils"

	"github.com/gin-gonic/gin"
)

// GetStudentsHandler godoc
//
//	@Summary		Get all Students
//	@Description	Retrieves a list of all students
//	@Tags			students
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/students/ [get]
func GetStudentsHandler(c *gin.Context) {
	students, err := service.ListStudents(c.Request.Context())
	if err != nil {
		utils.ResponseBody(c, 500, "Failed to retrieve students", nil)
		return
	}

	if len(students) == 0 {
		utils.ResponseBody(c, http.StatusOK, "No students found", []model.Student{})
		return
	}

	utils.ResponseBody(c, 200, "Students retrieved successfully", students)
}

// CreateStudentHandler godoc
//
//	@Summary		Create a new student
//	@Description	Creates a new student with the provided details
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			student	body	model.Student	true	"Student details"
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/api/v1/students/ [post]
func CreateStudentHandler(c *gin.Context) {
	var student model.Student

	// Validate input
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.ResponseBody(c, 400, "Invalid request payload: "+err.Error(), nil)
		return
	}

	// Insert into DB
	if err := service.CreateStudent(c.Request.Context(), &student); err != nil {
		utils.ResponseBody(c, 500, "Failed to create student", nil)
		return
	}

	utils.ResponseBody(c, 201, "Student created successfully", nil)
}
