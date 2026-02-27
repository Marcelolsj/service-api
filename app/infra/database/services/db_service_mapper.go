package database_services

import (
	"service-api/domain/models"
)

func MapDBServiceToServiceModel(dbService DBService) models.ServiceModel {
	return models.ServiceModel{
		ID:        dbService.ID,
		Descricao: dbService.Descricao,
		Valor:     dbService.Valor,
	}
}

func MapDBServicesToServicesModel(dbServices []DBService) []models.ServiceModel {
	services := make([]models.ServiceModel, len(dbServices))
	for i, dbService := range dbServices {
		services[i] = MapDBServiceToServiceModel(dbService)
	}
	return services
}
