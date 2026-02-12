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
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (this *UserController) UpdateUser(ctx *fiber.Ctx) error {
	var req dto_users.UpdateUserRequestDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	id := ctx.Params("id")

	response, err := this.UserService.UpdateUser(ctx.UserContext(), req, id)
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (this *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := this.UserService.DeleteUser(ctx.UserContext(), id); err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (this *UserController) GetUsers(ctx *fiber.Ctx) error {
	users, err := this.UserService.GetUsers(ctx.UserContext())
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (this *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := this.UserService.GetUser(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: *userService,
	}
}
