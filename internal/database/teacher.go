package database

import (
	"context"
)

type Teacher struct {
	Email string
}

func (db *Database) CreateTeacher(email string) (Teacher, error) {
	row := db.Database.QueryRowContext(context.TODO(), `INSERT INTO Teachers VALUES (?) RETURNING *;`, email)
	var createdTeacher Teacher
	err := row.Scan(&createdTeacher.Email)
	if err != nil {
		return Teacher{}, ConflictingTeachersEntry
	}

	return createdTeacher, nil
}
