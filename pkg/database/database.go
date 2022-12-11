package database 

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ReadDatabaseUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")

	if len(databaseUrl) == 0 {
		log.Fatalln("Failed to load 'DATABASE_URL', are you sure it's defined?")
	}

	return databaseUrl
}

func OpenDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", ReadDatabaseUrl())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
