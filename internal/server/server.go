package server

import (
	"github.com/gin-gonic/gin"
	"github.com/woojiahao/onecv_assignment/internal/database"
)

func Start(db *database.Database) {
	gin.Default().Run()
}
