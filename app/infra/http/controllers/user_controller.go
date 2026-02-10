package controllers

import (
	dto_users "service-api/application/dtos/users"
	"service-api/application/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService services.UserService
}

func (this *UserController) CreateUser(ctx *fiber.Ctx) error {
	var req dto_users.CreateUserRequestDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := this.UserService.CreateUser(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: *userService,
	}
}
