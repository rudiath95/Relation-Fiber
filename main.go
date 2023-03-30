package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
	"github.com/rudiath95/RelationFiber/routes"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
	ini.SyncDatabases()
}

func main() {
	app := fiber.New()

	routes.RouteIndex(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"test":  12,
			"test2": "aaa",
		})
	})

	app.Listen(":" + os.Getenv("PORT"))
}
