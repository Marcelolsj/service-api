package database_logs

import (
	"context"
	"service-api/domain/models"
	"service-api/domain/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DbLogRepository struct {
	db *sqlx.DB
}

func (this *DbLogRepository) CreateLog(ctx context.Context, log *models.LogModel) error {
	id := uuid.New().String()

	_, err := this.db.ExecContext(ctx,
		`INSERT INTO logs (id, id_recurso, log, criado_por)
		VALUES(?, ?, ?, ?) `,
		id,
		log.IDRecurso,
		log.Log,
		log.CriadoPor,
	)
	if err != nil {
		return err
	}

	log.ID = id

	return nil
}

func (this *DbLogRepository) GetLogs(ctx context.Context, resourceId string, startDate time.Time) ([]models.LogModel, error) {
	var logs []DBLog

	query := `
		SELECT id, id_recurso, log, criado_por, criado_em
		FROM logs
		WHERE (id_recurso = ? or ? = '') AND criado_em >= ?
	`

	err := this.db.Select(&logs, query, resourceId, resourceId, startDate)
	if err != nil {
		return nil, err
	}

	logModels := MapDBLogsToLogModels(logs)
	return logModels, nil
}

func NewDbLogRepository(db *sqlx.DB) repository.LogRepository {
	return &DbLogRepository{
		db: db,
	}
}
