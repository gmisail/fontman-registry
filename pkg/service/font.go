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
func GetFont(client *sqlx.DB, id uuid.UUID) (*model.Font, error) {
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

func DeleteFont() {

}
