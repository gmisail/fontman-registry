package service

import (
	"fontman/registry/pkg/model"
	"fontman/registry/pkg/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

/**
 * Get a font & its styles by a family ID
 */
func GetFontById(client *sqlx.DB, id uuid.UUID) (*model.Font, error) {
	font, fontErr := repository.GetFontById(client, id)

	if fontErr != nil {
		return nil, fontErr
	}

	styles, styleErr := repository.GetStylesByFamilyId(client, id)

	if styleErr != nil {
		return nil, styleErr
	}

	font.Styles = styles

	return font, nil
}

func GetFontByName(client *sqlx.DB, name string) ([]*model.Font, error) {
	return repository.GetFontsByName(client, name)
}

func UploadFont(client *sqlx.DB, payload model.FamilyUploadRequest) error {
	familyId, familyErr := repository.InsertFontFamily(
		client, payload.Name, payload.License, payload.Creator,
	)

	if familyErr != nil {
		return familyErr
	}

	return repository.InsertFontStyles(client, familyId, payload.Styles)
}

func DeleteFont() {

}
