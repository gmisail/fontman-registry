package repository

import (
	"fontman/registry/pkg/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

/**
 * Get information about a specific font family. Does not include font styles.
 */
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

/**
 * Get all styles for a font family.
 */
func GetStylesByFamilyId(client *sqlx.DB, familyId uuid.UUID) ([]*model.FontStyle, error) {
	var styles []*model.FontStyle

	err := client.Select(&styles, `
		SELECT *
		FROM FontStyle f
		WHERE f.family = ?
		LIMIT 1
	`, familyId)

	if err != nil {
		return nil, err
	}

	return styles, nil
}
