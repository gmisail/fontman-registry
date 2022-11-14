package model

import "github.com/google/uuid"

type Font struct {
	Id      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	License string    `db:"license"`
	Creator string    `db:"creator"`
}
