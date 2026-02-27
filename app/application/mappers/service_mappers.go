package mappers

import (
	dtos_services "service-api/application/dtos/services"
	"service-api/domain/models"
)

func MapCreateServiceDTOToService(createServiceDto dtos_services.CreateServiceRequestDTO) models.ServiceModel {
	return models.ServiceModel{
		Descricao: createServiceDto.Descricao,
		Valor:     createServiceDto.Valor,
	}
}

func MapUpdateServiceDTOToService(updateServiceDto dtos_services.UpdateServiceRequestDTO, id string) models.ServiceModel {
	return models.ServiceModel{
		ID:        id,
		Descricao: updateServiceDto.Descricao,
		Valor:     updateServiceDto.Valor,
	}
}

func MapServiceToServiceResponseDTO(service models.ServiceModel) *dtos_services.ServiceResponseDTO {
	return &dtos_services.ServiceResponseDTO{
		ID:        service.ID,
		Descricao: service.Descricao,
		Valor:     service.Valor,
	}
}

func MapServicesToServiceResponsesDTO(services []models.ServiceModel) []dtos_services.ServiceResponseDTO {
	responses := make([]dtos_services.ServiceResponseDTO, len(services))
	for i, service := range services {
		responses[i] = *MapServiceToServiceResponseDTO(service)
	}
	return responses
}
