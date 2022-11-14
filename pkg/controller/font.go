package controller

import (
	"fontman/registry/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FontController struct {
	Controller
}

func (f *FontController) Setup(app fiber.Router, db *sqlx.DB) {
	font := app.Group("/font")

	font.Get("/:id", func(ctx *fiber.Ctx) error {
		return GetFont(ctx, db)
	})

	font.Post("/", func(ctx *fiber.Ctx) error {
		return UploadFont(ctx, db)
	})
}

func GetFont(ctx *fiber.Ctx, client *sqlx.DB) error {
	id := ctx.Params("id")
	fontId := uuid.MustParse(id)

	font, err := service.GetFont(client, fontId)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"id":   font.Id,
		"name": font.Name,
	})
}

func UploadFont(ctx *fiber.Ctx, client *sqlx.DB) error {
	return ctx.SendString("upload font")
}
