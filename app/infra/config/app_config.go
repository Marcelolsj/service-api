package config

import (
	"service-api/application/services"
	"service-api/domain/use_case"
	database_users "service-api/infra/database/users"
	"service-api/infra/http/controllers"
	"service-api/infra/http/routers"

	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	routers.SetupRoutes(app,
		newUserController(),
	)

	return app
}

func newUserController() *controllers.UserController {
	user_repository := database_users.NewDbUserRepository()
	user_use_case := use_case.NewUserUseCase(user_repository)
	userService := services.NewUserService(user_use_case)
	controller := controllers.NewUserController(userService)
	return controller
}
