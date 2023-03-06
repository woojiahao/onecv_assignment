package database

import (
	"context"
)

type Student struct {
	Email       string
	IsSuspended bool
}

func (db *Database) CreateStudent(email string) (Student, error) {
	row := db.Database.QueryRowContext(context.TODO(), `INSERT INTO Students (email) VALUES (?) RETURNING *;`, email)
	var createdStudent Student
	err := row.Scan(&createdStudent.Email)
	if err != nil {
		return Student{}, ConflictingStudentsEntry
	}

	return createdStudent, nil
}
