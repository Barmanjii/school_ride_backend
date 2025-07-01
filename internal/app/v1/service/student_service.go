package service

import (
	"context"
	"errors"
	"school_ride_backend/internal/app/v1/config"
	"school_ride_backend/internal/app/v1/model"
)

func CreateStudent(ctx context.Context, student *model.Student) error {
	sql := `INSERT INTO student_info (name, age, gender, class, school_id) VALUES ($1, $2, $3, $4, $5)`

	cmdTag, err := config.DB.Exec(ctx, sql, student.Name, student.Age, student.Gender, student.Class, student.SchoolID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("no rows created for student")
	}
	return nil
}

func ListStudents(ctx context.Context) ([]*model.StudentInDatabase, error) {
	sql := `SELECT id, name, age, gender, class, school_id FROM student_info`
	rows, err := config.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*model.StudentInDatabase
	for rows.Next() {
		var s model.StudentInDatabase
		if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Gender, &s.Class, &s.SchoolID); err != nil {
			return nil, err
		}
		students = append(students, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
