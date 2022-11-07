package model

import "github.com/google/uuid"

type FontStyle struct {
	Id     uuid.UUID
	Family string
	Style  string
	URL    string
}
