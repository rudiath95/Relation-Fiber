package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"test":  12,
			"test2": "sss",
		})
	})

	app.Listen(":3000")
}
