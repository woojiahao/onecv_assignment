package main

import (
	"github.com/woojiahao/onecv_assignment/internal/database"
	"github.com/woojiahao/onecv_assignment/internal/server"
)

func main() {
	databaseConfiguration := database.LoadConfiguration()
	db := database.Connect(databaseConfiguration)
	server.Start(db)
}
