package controllers

import (
	dto_services "service-api/application/dtos/services"
	"service-api/application/services"

	"github.com/gofiber/fiber/v2"
)

type ServiceController struct {
	ServiceService services.ServiceService
}

func (this *ServiceController) CreateService(ctx *fiber.Ctx) error {
	var req dto_services.CreateServiceRequestDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := this.ServiceService.CreateService(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (this *ServiceController) UpdateService(ctx *fiber.Ctx) error {
	var req dto_services.UpdateServiceRequestDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	id := ctx.Params("id")

	response, err := this.ServiceService.UpdateService(ctx.UserContext(), req, id)
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (this *ServiceController) DeleteService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := this.ServiceService.DeleteService(ctx.UserContext(), id); err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (this *ServiceController) GetServices(ctx *fiber.Ctx) error {
	services, err := this.ServiceService.GetServices(ctx.UserContext())
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(services)
}

func (this *ServiceController) GetService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	service, err := this.ServiceService.GetService(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{"error": err.Error()})
	}
	if service == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.Status(fiber.StatusOK).JSON(service)
}

func NewServiceController(serviceService *services.ServiceService) *ServiceController {
	return &ServiceController{
		ServiceService: *serviceService,
	}
}
