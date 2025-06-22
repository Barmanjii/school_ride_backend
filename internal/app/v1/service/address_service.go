package service

import (
	"context"
	"errors"
	"school_ride_backend/internal/app/v1/config"
	"school_ride_backend/internal/app/v1/model"
)

func CreateAddress(ctx context.Context, address *model.Address) error {
	query := `
		INSERT INTO addresses (flat, street, landmark, city, state, pincode, country)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	cmdTag, err := config.DB.Exec(ctx, query, address.Flat, address.Street, address.Landmark, address.City, address.State, address.Pincode, address.Country)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no rows created for address")
	}
	return nil
}
func GetAddress(ctx context.Context, id string) (*model.Address, error) {
	return nil, nil
}
func UpdateAddress(ctx context.Context, address *model.Address) error {
	return nil
}
func DeleteAddress(ctx context.Context, id string) error {
	return nil
}
func ListAddresses(ctx context.Context) ([]*model.Address, error) {
	return nil, nil
}
