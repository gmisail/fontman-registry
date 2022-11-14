package service

import (
	"fontman/registry/pkg/model"
	"fontman/registry/pkg/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func GetFont(client *sqlx.DB, id uuid.UUID) (*model.Font, error) {
	return repository.GetFontById(client, id)
}

func DeleteFont() {

}

func GetStylesFromFont(font string, exclude []string) {

}
