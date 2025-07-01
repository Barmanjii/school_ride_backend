package model

import "github.com/google/uuid"

type Student struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Gender   uuid.UUID `json:"gender"`
	Class    uuid.UUID `json:"class"`
	SchoolID uuid.UUID `json:"school_id"`
}

type StudentInDatabase struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Gender   uuid.UUID `json:"gender"`
	Class    uuid.UUID `json:"class"`
	SchoolID uuid.UUID `json:"school_id"`
}
