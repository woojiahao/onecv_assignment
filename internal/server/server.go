package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woojiahao/onecv_assignment/internal/database"
	"github.com/woojiahao/onecv_assignment/internal/utility"
	"net/http"
)

func Start(db *database.Database) {
	engine := gin.Default()

	engine.POST("/api/teachers", func(context *gin.Context) {
		var createTeacher CreateTeacherDto
		if err := context.ShouldBindJSON(&createTeacher); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that email field contains email"})
			return
		}

		err := db.CreateTeacher(createTeacher.Email)
		if err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for teachers"})
			return
		}

		context.Status(http.StatusNoContent)
	})

	engine.POST("/api/students", func(context *gin.Context) {
		var createStudent CreateStudentDto
		if err := context.ShouldBindJSON(&createStudent); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input, ensure that email field contains email"})
			return
		}

		err := db.CreateStudent(createStudent.Email)
		if err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Duplicate email used for students"})
			return
		}

		context.Status(http.StatusNoContent)
	})

	engine.POST("/api/register", func(context *gin.Context) {
		var registerStudents RegisterStudentsDto
		if err := context.ShouldBindJSON(&registerStudents); err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that all field values are emails"})
			return
		}
		err := db.RegisterStudents(registerStudents.Teacher, registerStudents.Students)
		if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
			return
		}
		context.Status(http.StatusNoContent)
	})

	engine.GET("/api/commonstudents", func(context *gin.Context) {
		teacherEmails := context.QueryArray("teacher")
		students, err := db.GetStudents(teacherEmails...)
		if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"students": utility.Map(students, func(s database.Student) string {
				return s.Email
			}),
		})
	})

	engine.POST("/api/suspend", func(context *gin.Context) {
		var suspend SuspendDto
		if err := context.ShouldBindJSON(&suspend); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that all field values are emails"})
			return
		}
		err := db.Suspend(suspend.Student)
		if err != nil && err == database.NoStudentFound {
			context.JSON(http.StatusNotFound, ErrorResponse{"Student not found"})
			return
		} else if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
			return
		}

		context.Status(http.StatusNoContent)
	})

	engine.POST("/api/retrievefornotifications", func(context *gin.Context) {
		var retrieveForNotifications RetrieveForNotificationsDto
		if err := context.ShouldBindJSON(&retrieveForNotifications); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input: ensure that teacher field value is an email"})
			return
		}
		students, err := db.GetNotifiableStudents(retrieveForNotifications.Teacher, retrieveForNotifications.Notification)
		if err != nil {
			context.JSON(http.StatusInternalServerError, ErrorResponse{"Internal Server Error"})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"recipients": utility.Map(students, func(s database.Student) string {
				return s.Email
			}),
		})
	})

	engine.Run()
}
