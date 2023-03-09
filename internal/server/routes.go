package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woojiahao/onecv_assignment/internal/database"
	"github.com/woojiahao/onecv_assignment/internal/utility"
	"net/http"
)

func CreateTeacher(context *gin.Context, d *database.Database) {
	var createTeacher CreateTeacherDto
	if err := context.ShouldBindJSON(&createTeacher); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that email field contains email"})
		return
	}

	err := d.CreateTeacher(createTeacher.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for teachers"})
		return
	}

	context.Status(http.StatusNoContent)
}

func CreateStudent(context *gin.Context, d *database.Database) {
	var createStudent CreateStudentDto
	if err := context.ShouldBindJSON(&createStudent); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that email field contains email"})
		return
	}

	err := d.CreateStudent(createStudent.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for students"})
		return
	}

	context.Status(http.StatusNoContent)
}

func RegisterStudents(context *gin.Context, d *database.Database) {
	var registerStudents RegisterStudentsDto
	if err := context.ShouldBindJSON(&registerStudents); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that all field values are emails"})
		return
	}
	err := d.RegisterStudents(registerStudents.Teacher, registerStudents.Students)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
		return
	}
	context.Status(http.StatusNoContent)
}

func GetCommonStudents(context *gin.Context, d *database.Database) {
	teacherEmails := context.QueryArray("teacher")
	if len(teacherEmails) == 0 {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Provide at least 1 \"teacher\" query parameter"})
		return
	}
	students, err := d.GetCommonStudents(teacherEmails...)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"students": utility.Map(students, func(s database.Student) string {
			return s.Email
		}),
	})
}

func SuspendStudent(context *gin.Context, d *database.Database) {
	var suspend SuspendDto
	if err := context.ShouldBindJSON(&suspend); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that all field values are emails"})
		return
	}
	err := d.Suspend(suspend.Student)
	if err != nil && err == database.NoStudentFound {
		context.JSON(http.StatusNotFound, ErrorResponse{"Student not found"})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
		return
	}

	context.Status(http.StatusNoContent)
}

func GetNotifiableStudents(context *gin.Context, d *database.Database) {
	var retrieveForNotifications RetrieveForNotificationsDto
	if err := context.ShouldBindJSON(&retrieveForNotifications); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that teacher field value is an email"})
		return
	}
	students, err := d.GetNotifiableStudents(retrieveForNotifications.Teacher, retrieveForNotifications.Notification)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"recipients": utility.Map(students, func(s database.Student) string {
			return s.Email
		}),
	})
}
