package routers

import (
	"service-api/infra/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupServiceRoutes(app *fiber.App, serviceController *controllers.ServiceController) {
	app.Post("/services", serviceController.CreateService)
	app.Put("/services/:id", serviceController.UpdateService)
	app.Delete("/services/:id", serviceController.DeleteService)
	app.Get("/services", serviceController.GetServices)
	app.Get("/services/:id", serviceController.GetService)
}
