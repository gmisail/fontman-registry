package db

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gmisail/fontman/registry/ent"
	"github.com/gmisail/fontman/registry/ent/migrate"

	"log"
)

/*
	Create a database connection, setup migration configuration.
*/
func CreateConnection() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "registry.db?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer client.Close()

	ctx := context.Background()

	migrateErr := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

	return client
}
