package model

import "github.com/google/uuid"

type FontStyle struct {
	Id     uuid.UUID `db:"id"`
	Family string    `db:"family"`
	Style  string    `db:"style"`
	URL    string    `db:"url"`
}
