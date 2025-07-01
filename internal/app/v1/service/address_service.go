package service

import (
	"context"
	"errors"
	"school_ride_backend/internal/app/v1/config"
	"school_ride_backend/internal/app/v1/model"
)

func CreateAddress(ctx context.Context, address *model.Address) error {
	query := `
		INSERT INTO addresses (place_id, full_address, lat, lng)
		VALUES ($1, $2, $3, $4)`
	cmdTag, err := config.DB.Exec(ctx, query, address.PlaceID, address.FullAddress, address.Latitude, address.Longitude)
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
	query := `
		SELECT place_id, full_address, lat, lng
		FROM addresses`
	rows, err := config.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*model.Address
	for rows.Next() {
		var address model.Address
		if err := rows.Scan(&address.PlaceID, &address.FullAddress, &address.Latitude, &address.Longitude); err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}
