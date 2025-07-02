package service

import (
	"context"
	"errors"
	"school_ride_backend/internal/app/v1/config"
	"school_ride_backend/internal/app/v1/model"
)

func CreateParent(ctx context.Context, parent *model.Parent) error {
	query := `
		INSERT INTO parent_info (full_name, age, mobile_number, email_address, emergency_contact_number, address_id, parent_relation)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	cmdTag, err := config.DB.Exec(ctx, query,
		parent.FullName,
		parent.Age,
		parent.MobileNumber,
		parent.EmailAddress,
		parent.EmergencyContactNumber,
		parent.AddressID,
		parent.RelationToChild,
	)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no rows created for parent")
	}
	return nil
}

func ListParents(ctx context.Context) ([]*model.ParentInDatabase, error) {
	query := `
		SELECT id, full_name, age, mobile_number, email_address, emergency_contact_number, address_id, parent_relation
		FROM parent_info`
	rows, err := config.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parents []*model.ParentInDatabase
	for rows.Next() {
		var parent model.ParentInDatabase
		if err := rows.Scan(
			&parent.ID,
			&parent.FullName,
			&parent.Age,
			&parent.MobileNumber,
			&parent.EmailAddress,
			&parent.EmergencyContactNumber,
			&parent.AddressID,
			&parent.RelationToChild,
		); err != nil {
			return nil, err
		}
		parents = append(parents, &parent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return parents, nil
}
