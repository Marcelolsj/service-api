package use_case

import (
	"context"
	"service-api/domain/models"
	"service-api/domain/repository"
	"time"
)

type ILogUseCase interface {
	CreateLog(ctx context.Context, resourceId string, log string) error
	GetLogs(ctx context.Context, resourceId string, startDate time.Time) ([]models.LogModel, error)
}

type LogUseCase struct {
	LogRepo repository.LogRepository
}

func (this LogUseCase) CreateLog(ctx context.Context, resourceId string, log string) error {
	logModel := models.LogModel{
		IDRecurso: resourceId,
		Log:       log,
		CriadoPor: "7dc6ac44-1859-4933-86ab-e41741715e51",
	}

	err := this.LogRepo.CreateLog(ctx, &logModel)
	if err != nil {
		return err
	}

	return nil
}

func (this LogUseCase) GetLogs(ctx context.Context, resourceId string, startDate time.Time) ([]models.LogModel, error) {
	logs, err := this.LogRepo.GetLogs(ctx, resourceId, startDate)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func NewLogUseCase(logRepo repository.LogRepository) *LogUseCase {
	return &LogUseCase{
		LogRepo: logRepo,
	}
}
