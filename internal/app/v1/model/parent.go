package model

import "github.com/google/uuid"

// ENUM RelationToChild represents the relationship of the parent to the child.
// It can be "Father", "Mother", or "Guardian".
// This is used to categorize the type of parent or guardian in the system.
// It is a string type with predefined constants for each relation.
type RelationToChild string

const (
	RelationFather   RelationToChild = "Father"
	RelationMother   RelationToChild = "Mother"
	RelationGuardian RelationToChild = "Guardian"
)

type Parent struct {
	FullName               string          `json:"full_name"`
	Age                    int             `json:"age"`
	MobileNumber           string          `json:"mobile_number"`
	EmailAddress           string          `json:"email_address"`
	EmergencyContactNumber string          `json:"emergency_contact_number"`
	AddressID              uuid.UUID       `json:"address_id"`
	RelationToChild        RelationToChild `json:"parent_relation"` // e.g., "Father", "Mother", "Guardian"
}

type ParentInDatabase struct {
	ID                     uuid.UUID       `json:"id"`
	FullName               string          `json:"full_name"`
	Age                    int             `json:"age"`
	MobileNumber           string          `json:"mobile_number"`
	EmailAddress           string          `json:"email_address"`
	EmergencyContactNumber string          `json:"emergency_contact_number"`
	AddressID              uuid.UUID       `json:"address_id"`
	RelationToChild        RelationToChild `json:"parent_relation"` // e.g., "Father", "Mother", "Guardian"
}
