package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	USER_DB     = "postgres"
	PASSWORD_DB = "postgres"
	NAME_DB     = "dc-a01"
)

func ConnectionDB() *sql.DB {
	connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", USER_DB, PASSWORD_DB, NAME_DB)
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Panic(("Error with DB CONNECTION"))
	}

	return db
}
