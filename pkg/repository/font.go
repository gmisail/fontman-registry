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
 * Return all font families that have the given name.
 */
func GetFontsByName(client *sqlx.DB, name string) ([]*model.Font, error) {
	var fonts []*model.Font

	err := client.Select(&fonts, `
		SELECT *
		FROM FontFamily f
		WHERE f.name = ?
	`, name)

	if err != nil {
		return nil, err
	}

	return fonts, nil
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
	`, familyId)

	if err != nil {
		return nil, err
	}

	return styles, nil
}

/**
 * Insert a font family.
 */
func InsertFontFamily(client *sqlx.DB, name, license, creator string) (uuid.UUID, error) {
	id := uuid.New()

	_, err := client.Exec(`
		INSERT INTO FontFamily (id, name, license, creator)
		VALUES (?, ?, ?, ?)
	`, id.String(), name, license, creator)

	return id, err
}

func InsertFontStyles(client *sqlx.DB, familyId uuid.UUID, styles []model.StyleUploadRequest) error {
	// insert each style in one shot
	transaction := client.MustBegin()

	// insert each style from the payload
	for _, style := range styles {
		id := uuid.New()

		transaction.MustExec(`
			INSERT INTO FontStyle (id, family, type, url)
			VALUES (?, ?, ?, ?)
		`, id.String(), familyId, style.Style, style.Url)
	}

	return transaction.Commit()
}
