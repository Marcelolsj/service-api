package services

import (
	"context"
	dtos_logs "service-api/application/dtos/logs"
	http_errors "service-api/application/errors"
	mappers "service-api/application/mappers"
	"service-api/domain/use_case"
	"time"
)

type LogService struct {
	LogUseCase use_case.ILogUseCase
}

func (this *LogService) CreateLog(ctx context.Context, resourceId string, log string) error {

	if err := this.LogUseCase.CreateLog(ctx, resourceId, log); err != nil {
		return http_errors.NewHttpError(err)
	}

	return nil
}

func (this *LogService) GetLogs(ctx context.Context, resourceId string, startDate time.Time) ([]dtos_logs.LogResponseDTO, *http_errors.HttpError) {
	logs, err := this.LogUseCase.GetLogs(ctx, resourceId, startDate)
	if err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapLogsToDTOs(logs), nil
}

func NewLogService(logUseCase use_case.ILogUseCase) *LogService {
	return &LogService{
		LogUseCase: logUseCase,
	}
}
