package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App, userController *controllers.UserController) {
	app.Post("/users", userController.CreateUser)
}
