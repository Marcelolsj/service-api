package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupInfraRoutes(app *fiber.App) {
	app.Get("/health", controllers.HandleHealth)
}
