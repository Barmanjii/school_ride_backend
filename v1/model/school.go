package model

import "github.com/google/uuid"

type School struct {
    ID        uuid.UUID `json:"id"`
    Name      string    `json:"name"`
    Code      string    `json:"code"`
    AddressID uuid.UUID `json:"address_id"`
}
