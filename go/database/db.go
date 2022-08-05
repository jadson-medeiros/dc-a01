package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

const (
	USER_DB     = "postgres"
	PASSWORD_DB = "postgres"
	HOST_DB     = "db"
	PORT_DB     = "5432"
)

func ConnectionDB() ksql.DB {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", USER_DB, PASSWORD_DB, HOST_DB, PORT_DB)

	ctx := context.Background()
	db, err := kpgx.New(ctx, connectionString, ksql.Config{})

	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}

	return db
}
