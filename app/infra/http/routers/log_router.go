package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupLogRoutes(app *fiber.App, logController *controllers.LogController) {
	app.Get("/logs", logController.GetLogs)
}
