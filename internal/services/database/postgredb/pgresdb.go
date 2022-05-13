package postgredb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Hand = Connect()

type Postgres struct {
	db *sql.DB
}

func Connect() *Postgres {
	uri := "host=127.0.0.1 port=5432 user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Printf("[Crash]\tPostgres Not Connected\tURI: '%s'\tError: '%s'\n", uri, err.Error())
		os.Exit(1)
	}
	return &Postgres{
		db: db,
	}
}
