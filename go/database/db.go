package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	USER_DB     = "postgres"
	PASSWORD_DB = "postgres"
	NAME_DB     = "postgres"
	HOST_DB     = "127.0.0.1"
	PORT_DB     = "5432"
)

func ConnectionDB() *sql.DB {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", USER_DB, PASSWORD_DB, HOST_DB, PORT_DB, NAME_DB)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Panic(("Error with DB CONNECTION"))
	}

	return db
}
