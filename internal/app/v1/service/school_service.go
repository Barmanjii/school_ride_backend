package service

import (
	"context"
	"errors"
	"school_ride_backend/internal/app/v1/config"
	"school_ride_backend/internal/app/v1/model"
)

func CreateSchool(ctx context.Context, school *model.School) error {

	sql := `INSERT INTO school (name, code, address_id) VALUES ($1, $2, $3)`

	cmdTag, err := config.DB.Exec(ctx, sql, school.Name, school.Code, school.AddressID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no rows created for school")
	}
	return nil
}

func GetAllSchools(ctx context.Context) ([]model.School, error) {
	sql := `SELECT id, name, code, address_id FROM school`

	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []model.School
	for rows.Next() {
		var school model.School
		if err := rows.Scan(&school.Name, &school.Code, &school.AddressID); err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil
}
