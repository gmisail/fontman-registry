package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserController struct {
	Controller
}

func (u *UserController) Setup(app fiber.Router, db *sqlx.DB) {
	fmt.Println("setup user controller")
}
