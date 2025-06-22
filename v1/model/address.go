package model

import "github.com/google/uuid"
type Address struct {
	ID      uuid.UUID `json:"id"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
}
