package controller

import (
	"fontman/registry/pkg/model"
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
		return GetFontById(ctx, db)
	})

	font.Get("/", func(ctx *fiber.Ctx) error {
		return GetFontsByParam(ctx, db)
	})

	font.Post("/", func(ctx *fiber.Ctx) error {
		return UploadFont(ctx, db)
	})
}

func GetFontById(ctx *fiber.Ctx, client *sqlx.DB) error {
	id := ctx.Params("id")
	fontId, idErr := uuid.Parse(id)

	if idErr != nil {
		return idErr
	}

	font, err := service.GetFontById(client, fontId)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"id":     font.Id,
		"name":   font.Name,
		"styles": font.Styles,
	})
}

func GetFontsByParam(ctx *fiber.Ctx, client *sqlx.DB) error {
	name := ctx.Query("name")

	// no name provided? Send back empty list
	if len(name) == 0 {
		return ctx.JSON(fiber.Map{"fonts": []*model.Font{}})
	}

	fonts, err := service.GetFontByName(client, name)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"fonts": fonts,
	})
}

/*
Expects request with structure:
	{
		"name": "Menlo",
		"license": "",
		"creator": "",
		"styles": [
			{ "style": "Regular", "url": "gmisail.me/regular.ttf" },
			{ "style": "Italic", "url": "gmisail.me/italic.ttf" },
			{ "style": "Bold", "url": "gmisail.me/bold.ttf" }
		]
	}
*/
func UploadFont(ctx *fiber.Ctx, client *sqlx.DB) error {
	payload := model.FamilyUploadRequest{}

	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if err := service.UploadFont(client, payload); err != nil {
		return err
	}

	return ctx.JSON(payload)
}
