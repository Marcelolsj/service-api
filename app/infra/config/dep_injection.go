package config

import (
	"service-api/application/services"
	"service-api/domain/use_case"
	"service-api/infra/database"
	database_users "service-api/infra/database/users"
	"service-api/infra/http/controllers"
)

func NewUserController() *controllers.UserController {
	user_repository := database_users.NewDbUserRepository(database.GetConnection())
	user_use_case := use_case.NewUserUseCase(user_repository)
	userService := services.NewUserService(user_use_case)
	controller := controllers.NewUserController(userService)
	return controller
}
