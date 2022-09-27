package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"

	"github.com/gmisail/fontman/registry/ent"
	"github.com/gmisail/fontman/registry/ent/migrate"

	"log"
)

/*
	Create a database connection, setup migration configuration.
*/
func CreateConnection() *ent.Client {
	databaseFile := os.Getenv("DATABASE_FILE")
	client, err := ent.Open(dialect.SQLite, fmt.Sprintf("%s?mode=memory&cache=shared&_fk=1", databaseFile))

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
