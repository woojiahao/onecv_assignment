package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     int
}

func LoadConfiguration() *Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env file")
	}

	port := os.Getenv("DATABASE_PORT")
	portValue, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Unable to parse port value as integer")
	}

	return &Configuration{
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		portValue,
	}
}

type Database struct {
	Configuration *Configuration
	Database      *sql.DB
}

func Connect(c *Configuration) *Database {
	connStr := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?ssl-mode=disabled", c.Username, c.Password, c.Host, c.Port, c.Name)
	if c.Password == "" {
		connStr = fmt.Sprintf("mysql://%s@tcp(%s:%d)/%s?ssl-mode=disabled", c.Username, c.Host, c.Port, c.Name)
	}
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database due to %s\n", err)
	}

	return &Database{c, db}
}
