package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Font struct {
	Id      uuid.UUID      `db:"id" json:"id"`
	Name    string         `db:"name" json:"name"`
	License sql.NullString `db:"license" json:"license"`
	Creator sql.NullString `db:"creator" json:"creator"`
	Styles  []*FontStyle   `json:"styles"`
}
