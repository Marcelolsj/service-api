package services

import (
	"context"
	dtos_services "service-api/application/dtos/services"
	http_errors "service-api/application/errors"
	"service-api/application/mappers"
	"service-api/domain/use_case"
)

type ServiceService struct {
	ServiceUseCase use_case.IServiceUseCase
	LogService     *LogService
}

func (this *ServiceService) CreateService(ctx context.Context, createServiceDto dtos_services.CreateServiceRequestDTO) (*dtos_services.ServiceResponseDTO, *http_errors.HttpError) {
	service := mappers.MapCreateServiceDTOToService(createServiceDto)

	if err := this.ServiceUseCase.CreateService(ctx, &service); err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	this.LogService.CreateLog(ctx, service.ID, "Serviço criado")

	return mappers.MapServiceToServiceResponseDTO(service), nil
}

func (this *ServiceService) UpdateService(ctx context.Context, updateServiceDto dtos_services.UpdateServiceRequestDTO, id string) (*dtos_services.ServiceResponseDTO, *http_errors.HttpError) {
	service := mappers.MapUpdateServiceDTOToService(updateServiceDto, id)

	if err := this.ServiceUseCase.UpdateService(ctx, &service); err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	this.LogService.CreateLog(ctx, service.ID, "Serviço atualizado")

	return mappers.MapServiceToServiceResponseDTO(service), nil
}

func (this *ServiceService) DeleteService(ctx context.Context, id string) *http_errors.HttpError {
	if err := this.ServiceUseCase.DeleteService(ctx, id); err != nil {
		return http_errors.NewHttpError(err)
	}

	this.LogService.CreateLog(ctx, id, "Serviço deletado")

	return nil
}

func (this *ServiceService) GetServices(ctx context.Context) ([]dtos_services.ServiceResponseDTO, *http_errors.HttpError) {
	services, err := this.ServiceUseCase.GetServices(ctx)
	if err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapServicesToServiceResponsesDTO(services), nil
}

func (this *ServiceService) GetService(ctx context.Context, id string) (*dtos_services.ServiceResponseDTO, *http_errors.HttpError) {
	service, err := this.ServiceUseCase.GetService(ctx, id)
	if err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapServiceToServiceResponseDTO(*service), nil
}

func NewServiceService(serviceUseCase use_case.IServiceUseCase, logService *LogService) *ServiceService {
	return &ServiceService{
		ServiceUseCase: serviceUseCase,
		LogService:     logService,
	}
}
