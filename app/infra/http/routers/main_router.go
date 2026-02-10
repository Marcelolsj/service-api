package routers

import (
	"service-api/infra/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupInfraRoutes(app)
	setupUserRoutes(app, config.NewUserController())
}
