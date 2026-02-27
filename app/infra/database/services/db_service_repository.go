package database_services

import (
	"context"
	"database/sql"
	"service-api/domain/models"
	"service-api/domain/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DbServiceRepository struct {
	db *sqlx.DB
}

func (this *DbServiceRepository) CreateService(ctx context.Context, service *models.ServiceModel) error {
	id := uuid.New().String()

	_, err := this.db.ExecContext(ctx,
		`INSERT INTO servicos (id, descricao, valor, criado_por)
		VALUES(?, ?, ?, ?) `,
		id,
		service.Descricao,
		service.Valor,
		"7dc6ac44-1859-4933-86ab-e41741715e51",
	)
	if err != nil {
		return err
	}

	service.ID = id

	return nil
}

func (this *DbServiceRepository) UpdateService(ctx context.Context, service *models.ServiceModel) error {
	_, err := this.db.ExecContext(ctx,
		`update servicos
		 set descricao = ?, 
		 	 valor = ?
		 where id = ?`,
		service.Descricao,
		service.Valor,
		service.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (this *DbServiceRepository) DeleteService(ctx context.Context, id string) error {
	_, err := this.db.ExecContext(ctx,
		`delete from servicos where id = ?`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *DbServiceRepository) GetService(ctx context.Context, id string) (*models.ServiceModel, error) {
	var service DBService

	query := `
		SELECT id, descricao, valor
		FROM servicos
		WHERE id = ?
	`

	err := this.db.Get(&service, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	serviceModel := MapDBServiceToServiceModel(service)
	return &serviceModel, nil
}

func (this *DbServiceRepository) GetServices(ctx context.Context) ([]models.ServiceModel, error) {
	var services []DBService

	query := `
		SELECT id, descricao, valor
		FROM servicos
	`

	err := this.db.Select(&services, query)
	if err != nil {
		return nil, err
	}

	serviceModels := MapDBServicesToServicesModel(services)
	return serviceModels, nil
}

func NewDbServiceRepository(db *sqlx.DB) repository.ServiceRepository {
	return &DbServiceRepository{
		db: db,
	}
}
