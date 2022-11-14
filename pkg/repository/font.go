package repository

import (
	"fontman/registry/pkg/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func GetFontById(client *sqlx.DB, id uuid.UUID) (*model.Font, error) {
	var font model.Font

	err := client.Get(&font, `
		SELECT id, name 
		FROM FontFamily f
		WHERE f.id = ?
		LIMIT 1
	`, id)

	if err != nil {
		return nil, err
	}

	return &font, nil
}
