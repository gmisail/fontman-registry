package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type FontController struct {
	Controller
}

func (f *FontController) Setup(app *fiber.Router, db *sqlx.DB) {
	fmt.Println("setup font controller")
}
