package model

import "github.com/google/uuid"

type FontStyle struct {
	Id     uuid.UUID `db:"id"`
	Family string    `db:"family"`
	Type   string    `db:"type"`
	URL    string    `db:"url"`
}
