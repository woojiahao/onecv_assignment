package database

import (
	"context"
	"log"
)

type Teacher struct {
	Email string
}

func (db *Database) CreateTeacher(email string) Teacher {
	row := db.Database.QueryRowContext(
		context.TODO(), `
		INSERT INTO Teachers 
		VALUES (?) 
		ON CONFLICT 
		    DO NOTHING 
		RETURNING *;
		`,
		email)
	var createdTeacher Teacher
	err := row.Scan(&createdTeacher.Email)
	if err != nil {
		log.Printf("Unable to create new teacher due to %s\n", err)
	}

	return createdTeacher
}
