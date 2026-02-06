package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController) {
	setupInfraRoutes(app)
	setupUserRoutes(app, userController)
}
