package main

import (
	"github.com/gmisail/fontman/registry/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	client := db.CreateConnection()
	client.Debug()

	app.Listen(":8080")
}
