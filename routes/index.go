package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/controllers"
)

func RouteIndex(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)

	app.Get("/lockers", controllers.LockerGetAll)
	app.Post("/lockers", controllers.CreateLocker)

	app.Get("/posts", controllers.PostGetAll)
	app.Post("/posts", controllers.CreatePost)

}
