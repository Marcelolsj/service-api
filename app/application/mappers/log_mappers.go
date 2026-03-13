package mappers

import (
	"service-api/domain/models"

	dtos_logs "service-api/application/dtos/logs"
)

func MapLogsToDTOs(logs []models.LogModel) []dtos_logs.LogResponseDTO {
	var logDTOs []dtos_logs.LogResponseDTO
	for _, log := range logs {
		logDTO := dtos_logs.LogResponseDTO{
			ID:        log.ID,
			Log:       log.Log,
			CriadoPor: log.CriadoPor,
			CriadoEm:  log.CriadoEm.Format("02/01/2006 15:04:05"),
		}
		logDTOs = append(logDTOs, logDTO)
	}
	return logDTOs
}
