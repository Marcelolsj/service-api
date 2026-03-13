package database_logs

import "service-api/domain/models"

func MapDBLogToLogModel(dbLog DBLog) models.LogModel {
	return models.LogModel{
		ID:        dbLog.ID,
		IDRecurso: dbLog.IDRecurso,
		Log:       dbLog.Log,
		CriadoPor: dbLog.CriadoPor,
		CriadoEm:  dbLog.CriadoEm,
	}
}

func MapDBLogsToLogModels(dbLogs []DBLog) []models.LogModel {
	logModels := make([]models.LogModel, len(dbLogs))
	for i, dbLog := range dbLogs {
		logModels[i] = MapDBLogToLogModel(dbLog)
	}
	return logModels
}
