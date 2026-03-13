package repository

import (
	"context"
	"service-api/domain/models"
	"time"
)

type LogRepository interface {
	CreateLog(ctx context.Context, log *models.LogModel) error
	GetLogs(ctx context.Context, resourceId string, startDate time.Time) ([]models.LogModel, error)
}
