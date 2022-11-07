package model

import "github.com/google/uuid"

type Font struct {
	Id         uuid.UUID
	Name       string
	Creator    string
	UploadedBy string // TODO: should be a user model
}
