package model

import "github.com/google/uuid"

type School struct {
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	AddressID uuid.UUID `json:"address_id"`
}
