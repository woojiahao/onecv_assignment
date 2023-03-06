package database

import (
	"context"
	"fmt"
	"github.com/woojiahao/onecv_assignment/internal/utility"
	"strings"
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

// TODO: Return string of names instead?
func (db *Database) GetStudents(teacherEmail string) ([]Student, error) {
	rows, err := db.Database.QueryContext(context.TODO(), `SELECT student_email FROM TeacherStudents WHERE teacher_email = ?;`, teacherEmail)
	if err != nil {
		return nil, DatabaseError
	}
	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Email)
		students = append(students, student)
	}
	return students, nil
}
