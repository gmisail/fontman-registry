package repository

import (
	"github.com/gmisail/fontman/registry/pkg/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func GetFontById(client *sqlx.DB, id uuid.UUID) model.Font {
	var font model.Font

	err := client.Get(&font, "SELECT * FROM place LIMIT 1")

	return font
}

func GetFontByFamily(client *sqlx.DB, family string) model.Font {

}
