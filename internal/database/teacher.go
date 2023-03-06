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

func (db *Database) RegisterStudents(teacherEmail string, studentEmails []string) error {
	var parameters []string
	var placeholders []string
	for _, studentEmail := range studentEmails {
		parameters = append(parameters, []string{teacherEmail, studentEmail}...)
		placeholders = append(placeholders, fmt.Sprintf("(?, ?)"))
	}
	query := fmt.Sprintf(
		`INSERT INTO TeacherStudents VALUES %s ON CONFLICT DO NOTHING;`,
		strings.Join(placeholders, ", "),
	)
	_, err := db.Database.ExecContext(context.TODO(), query, parameters)
	if err != nil {
		return DatabaseError
	}

	return nil
}
