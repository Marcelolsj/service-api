package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App, userController *controllers.UserController) {
	app.Post("/users", userController.CreateUser)
	app.Put("/users/:id", userController.UpdateUser)
	app.Delete("/users/:id", userController.DeleteUser)
	app.Get("/users", userController.GetUsers)
	app.Get("/users/:id", userController.GetUser)
}
