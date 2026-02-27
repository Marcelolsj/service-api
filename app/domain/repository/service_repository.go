package repository

import (
	"context"
	"service-api/domain/models"
)

type ServiceRepository interface {
	CreateService(ctx context.Context, service *models.ServiceModel) error
	UpdateService(ctx context.Context, service *models.ServiceModel) error
	DeleteService(ctx context.Context, id string) error
	GetService(ctx context.Context, id string) (*models.ServiceModel, error)
	GetServices(ctx context.Context) ([]models.ServiceModel, error)
}
