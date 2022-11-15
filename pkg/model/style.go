package model

import "github.com/google/uuid"

type FontStyle struct {
	Id     uuid.UUID `db:"id" json:"-"`
	Family string    `db:"family" json:"-"`
	Type   string    `db:"type" json:"type"`
	URL    string    `db:"url" json:"url"`
}
