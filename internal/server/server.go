package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/woojiahao/onecv_assignment/internal/database"
	"log"
)

type (
	HttpMethod string

	Server struct {
		Engine   *gin.Engine
		Database *database.Database
	}
)

const (
	Post   HttpMethod = "POST"
	Get    HttpMethod = "GET"
	Put    HttpMethod = "PUT"
	Delete HttpMethod = "DELETE"
)

func (s *Server) Register(method HttpMethod, endpoint string, fn func(*gin.Context, *database.Database)) {
	s.Engine.Handle(string(method), endpoint, func(context *gin.Context) {
		fn(context, s.Database)
	})
}

func (s *Server) Cors() {
	s.Engine.Use(cors.Default())
}

func (s *Server) Start() {
	err := s.Engine.Run()
	if err != nil {
		log.Fatalf("Failed to start server")
	}
}

func Start(db *database.Database) {
	engine := gin.Default()

	server := Server{engine, db}
	server.Cors()
	server.Register(Post, "/api/teachers", CreateTeacher)
	server.Register(Post, "/api/students", CreateStudent)
	server.Register(Post, "/api/register", RegisterStudents)
	server.Register(Get, "/api/commonstudents", GetCommonStudents)
	server.Register(Post, "/api/suspend", SuspendStudent)
	server.Register(Post, "/api/retrievefornotifications", GetNotifiableStudents)

	server.Start()
}
