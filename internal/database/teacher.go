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

func (db *Database) CreateTeacher(email string) error {
	_, err := db.Database.ExecContext(context.TODO(), `INSERT INTO Teachers VALUES (?);`, email)
	if err != nil {
		return ConflictingTeachersEntry
	}

	return nil
}

func (db *Database) RegisterStudents(teacherEmail string, studentEmails []string) error {
	var parameters []any
	var placeholders []string
	for _, studentEmail := range studentEmails {
		parameters = append(parameters, []any{teacherEmail, studentEmail}...)
		placeholders = append(placeholders, fmt.Sprintf("(?, ?)"))
	}
	query := fmt.Sprintf(
		`INSERT IGNORE INTO TeacherStudents VALUES %s;`,
		strings.Join(placeholders, ", "),
	)
	_, err := db.Database.ExecContext(context.TODO(), query, parameters...)
	if err != nil {
		return DatabaseError
	}

	return nil
}

func (db *Database) GetCommonStudents(teacherEmails ...string) ([]Student, error) {
	var queries []string
	for i := 0; i < len(teacherEmails); i++ {
		queries = append(queries, "SELECT student_email FROM TeacherStudents WHERE teacher_email = ?")
	}
	query := strings.Join(queries, " INTERSECT ")
	rows, err := db.Database.QueryContext(
		context.TODO(),
		query,
		utility.Map(teacherEmails, func(email string) any {
			return email
		})...,
	)
	if err != nil {
		return nil, DatabaseError
	}
	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Email)
		if err != nil {
			return nil, DatabaseError
		}
		students = append(students, student)
	}
	return students, nil
}

func (db *Database) Suspend(studentEmail string) error {
	res, err := db.Database.ExecContext(context.TODO(), `UPDATE Students SET is_suspended = TRUE WHERE Students.email = ?;`, studentEmail)
	if err != nil {
		return DatabaseError
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return DatabaseError
	}
	if affected < 1 {
		return NoStudentFound
	}

	return nil
}

func (db *Database) GetNotifiableStudents(teacherEmail, notification string) ([]Student, error) {
	parameters := []any{teacherEmail}
	mentions := utility.GetMentionsFromNotification(notification)
	parameters = append(parameters, utility.Map(mentions, func(s string) any {
		return s
	})...)
	query := fmt.Sprintf(`SELECT Students.email, Students.is_suspended
		FROM Students 
			LEFT JOIN TeacherStudents ON student_email = email 
		WHERE (teacher_email = ? OR Students.email IN (%s)) AND NOT is_suspended
		`, utility.Repeat("?", len(mentions), ", "))
	if len(parameters) == 1 {
		query = `SELECT Students.email, Students.is_suspended
		FROM Students 
			LEFT JOIN TeacherStudents ON student_email = email 
		WHERE (teacher_email = ?) AND NOT is_suspended
		`
	}
	rows, err := db.Database.QueryContext(context.TODO(), query, parameters...)
	if err != nil {
		return nil, DatabaseError
	}
	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Email, &student.IsSuspended)
		if err != nil {
			return nil, DatabaseError
		}
		students = append(students, student)
	}
	return students, nil
}
