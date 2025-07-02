package model

import "github.com/google/uuid"

type Address struct {
	PlaceID     string `json:"place_id"`
	FullAddress string `json:"full_address"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lng"`
}

type AddressInDatabase struct {
	ID          uuid.UUID `json:"id"`
	PlaceID     string    `json:"place_id"`
	FullAddress string    `json:"full_address"`
	Latitude    string    `json:"lat"`
	Longitude   string    `json:"lng"`
}
