package server

import (
	"github.com/gin-gonic/gin"
	"github.com/woojiahao/onecv_assignment/internal/database"
	"github.com/woojiahao/onecv_assignment/internal/utility"
	"net/http"
)

var InternalServerError = ErrorResponse{"Internal Server Error"}

func CreateTeacher(context *gin.Context, d *database.Database) {
	var createTeacher CreateTeacherDto
	if err := context.ShouldBindJSON(&createTeacher); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that 'email' field is present and is a valid email"})
		return
	}

	err := d.CreateTeacher(createTeacher.Email)
	if err == database.InvalidEmail {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid email given"})
		return
	} else if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for teacher"})
		return
	}

	context.Status(http.StatusNoContent)
}

func CreateStudent(context *gin.Context, d *database.Database) {
	var createStudent CreateStudentDto
	if err := context.ShouldBindJSON(&createStudent); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that 'email' field is present and is a valid email"})
		return
	}

	err := d.CreateStudent(createStudent.Email)
	if err == database.InvalidEmail {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid email given"})
		return
	} else if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for student"})
		return
	}

	context.Status(http.StatusNoContent)
}

func RegisterStudents(context *gin.Context, d *database.Database) {
	var registerStudents RegisterStudentsDto
	if err := context.ShouldBindJSON(&registerStudents); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that 'teacher' and 'students' fields are present and all values are valid emails"})
		return
	}
	err := d.RegisterStudents(registerStudents.Teacher, registerStudents.Students)
	if err != nil {
		context.JSON(http.StatusInternalServerError, InternalServerError)
		return
	}
	context.Status(http.StatusNoContent)
}

func GetCommonStudents(context *gin.Context, d *database.Database) {
	teacherEmails := context.QueryArray("teacher")
	if len(teacherEmails) == 0 {
		context.JSON(http.StatusBadRequest, ErrorResponse{`Provide at least 1 'teacher' query parameter`})
		return
	}
	students, err := d.GetCommonStudents(teacherEmails...)
	if err != nil {
		context.JSON(http.StatusInternalServerError, InternalServerError)
		return
	}

	result := make([]string, 0)
	if len(students) > 0 {
		result = utility.Map(students, func(s database.Student) string {
			return s.Email
		})
	}
	context.JSON(http.StatusOK, gin.H{"students": result})
}

func SuspendStudent(context *gin.Context, d *database.Database) {
	var suspend SuspendDto
	if err := context.ShouldBindJSON(&suspend); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that field 'student' is present and value is an email"})
		return
	}
	err := d.Suspend(suspend.Student)
	if err != nil && err == database.NoStudentFound {
		context.JSON(http.StatusNotFound, ErrorResponse{"Student not found or already suspended"})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, InternalServerError)
		return
	}

	context.Status(http.StatusNoContent)
}

func GetNotifiableStudents(context *gin.Context, d *database.Database) {
	var retrieveForNotifications RetrieveForNotificationsDto
	if err := context.ShouldBindJSON(&retrieveForNotifications); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that 'teacher' and 'notification' fields are present and teacher field value is an email"})
		return
	}
	students, err := d.GetNotifiableStudents(retrieveForNotifications.Teacher, retrieveForNotifications.Notification)
	if err != nil {
		context.JSON(http.StatusInternalServerError, InternalServerError)
		return
	}

	result := make([]string, 0)
	if len(students) > 0 {
		result = utility.Map(students, func(s database.Student) string {
			return s.Email
		})
	}
	context.JSON(http.StatusOK, gin.H{"recipients": result})
}
