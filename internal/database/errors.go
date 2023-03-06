package database

import "errors"

var (
	ConflictingTeachersEntry = errors.New("conflicting teacher with same email")
	ConflictingStudentsEntry = errors.New("conflicting student with same email")
)
