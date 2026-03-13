package use_case

import (
	"context"
	domain_errors "service-api/domain/domain_errors"
	"service-api/domain/models"
	"service-api/domain/repository"
)

type IServiceUseCase interface {
	CreateService(ctx context.Context, service *models.ServiceModel) error
	UpdateService(ctx context.Context, service *models.ServiceModel) error
	DeleteService(ctx context.Context, id string) error
	GetService(ctx context.Context, id string) (*models.ServiceModel, error)
	GetServices(ctx context.Context) ([]models.ServiceModel, error)
}

type ServiceUseCase struct {
	ServiceRepo repository.ServiceRepository
}

func (this ServiceUseCase) CreateService(ctx context.Context, service *models.ServiceModel) error {
	err := this.validateService(service)
	if err != nil {
		return err
	}

	err = this.ServiceRepo.CreateService(ctx, service)
	if err != nil {
		return err
	}

	updatedService, err := this.GetService(ctx, service.ID)
	if updatedService != nil {
		*service = *updatedService
	} else {
		return domain_errors.NewBusinessErrorMsg("Serviço não encontrado após atualização")
	}

	return err
}

func (this ServiceUseCase) UpdateService(ctx context.Context, service *models.ServiceModel) error {
	err := this.validateService(service)
	if err != nil {
		return err
	}

	if service.ID == "" {
		return domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	err = this.ServiceRepo.UpdateService(ctx, service)
	if err != nil {
		return err
	}

	updatedService, err := this.GetService(ctx, service.ID)
	if updatedService != nil {
		*service = *updatedService
	} else {
		return domain_errors.NewBusinessErrorMsg("Serviço não encontrado após atualização")
	}

	return err
}

func (this ServiceUseCase) DeleteService(ctx context.Context, id string) error {
	if id == "" {
		return domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	err := this.ServiceRepo.DeleteService(ctx, id)
	if err != nil {
		return err
	}

	return err
}

func (this ServiceUseCase) GetService(ctx context.Context, id string) (*models.ServiceModel, error) {
	if id == "" {
		return nil, domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	service, err := this.ServiceRepo.GetService(ctx, id)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (this ServiceUseCase) GetServices(ctx context.Context) ([]models.ServiceModel, error) {
	services, err := this.ServiceRepo.GetServices(ctx)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s *ServiceUseCase) validateService(service *models.ServiceModel) error {
	if service.Descricao == "" {
		return domain_errors.NewBusinessErrorMsg("A descrição é obrigatória")
	}
	if service.Valor == 0 {
		return domain_errors.NewBusinessErrorMsg("O valor é obrigatório")
	}

	return nil
}

func NewServiceUseCase(serviceRepo repository.ServiceRepository) *ServiceUseCase {
	return &ServiceUseCase{
		ServiceRepo: serviceRepo,
	}
}
