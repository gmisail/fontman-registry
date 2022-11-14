package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Controller interface {
	Setup(app fiber.Router, db *sqlx.DB)
}
