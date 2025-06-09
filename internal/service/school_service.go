package service

import (
	"context"
	"errors"

	"school_ride_backend/internal/config"
	"school_ride_backend/internal/model"

	"github.com/google/uuid"
)

func CreateSchool(ctx context.Context, school *model.School) error {
    if school.ID == uuid.Nil {
        school.ID = uuid.New()
    }

    sql := `INSERT INTO school (id, name, code, address_id) VALUES ($1, $2, $3, $4)`

    cmdTag, err := config.DB.Exec(ctx, sql, school.ID, school.Name, school.Code, school.AddressID)
    if err != nil {
        return err
    }
    if cmdTag.RowsAffected() != 1 {
        return errors.New("no rows affected")
    }
    return nil
}
