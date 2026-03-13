package controllers

import (
	"service-api/application/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	LogService services.LogService
}

func (this *LogController) GetLogs(ctx *fiber.Ctx) error {
	resourceId := ctx.Query("resourceId")
	startDateStr := ctx.Query("startDate")

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	startDate, err := time.ParseInLocation("02/01/2006", startDateStr, loc)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid startDate format"})
	}

	logs, httpError := this.LogService.GetLogs(ctx.UserContext(), resourceId, startDate)
	if httpError != nil {
		return ctx.Status(httpError.Status).JSON(fiber.Map{"error": httpError.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func NewLogController(logService *services.LogService) *LogController {
	return &LogController{
		LogService: *logService,
	}
}
