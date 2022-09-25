package main

import (
	"github.com/gmisail/fontman/registry/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Listen(":8080")
}
