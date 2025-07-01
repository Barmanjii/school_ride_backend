package model

import "github.com/google/uuid"

type School struct {
	Name       string    `json:"name"`
	SchoolCode string    `json:"school_code"`
	AddressID  uuid.UUID `json:"address_id"`
}

type SchoolInDatabase struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SchoolCode string    `json:"school_code"`
	AddressID  uuid.UUID `json:"address_id"`
}
