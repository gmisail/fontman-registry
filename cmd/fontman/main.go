package main

import (
	"fontman/registry/pkg/database"
	"fontman/registry/pkg/controller"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2"
)

/*
	Given a list of controllers, initialize all the routes.
*/
func setupControllers(controllers []controller.Controller, app fiber.Router, db *sqlx.DB) {
	for _, controller := range controllers {
		controller.Setup(app, db)
	}
}

func main() {
	db := database.OpenDatabase() 

	app := fiber.New()
	api := app.Group("/api")

	setupControllers([]controller.Controller{
		&controller.FontController{},
		&controller.UserController{},
	}, api, db)

	port := os.Getenv("PORT")

	// default port to 8080 if PORT is undefined
	if len(port) == 0 {
		port = "8080"
	}

	app.Listen(":" + port)
}
