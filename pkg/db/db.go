package db

import (
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gmisail/fontman/registry/ent"
	"log"
)

func CreateConnection() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "registry.db")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer client.Close()

	/*
		run migrations...
	*/

	return client
}
