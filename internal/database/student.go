package database

import (
	"context"
)

type Student struct {
	Email       string
	IsSuspended bool
}

func (db *Database) CreateStudent(email string) error {
	_, err := db.Database.ExecContext(context.TODO(), `INSERT INTO Students (email) VALUES (?)`, email)
	if err != nil {
		return ConflictingStudentsEntry
	}

	return nil
}
