package dtos_services

type CreateServiceRequestDTO struct {
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

type UpdateServiceRequestDTO struct {
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

type ServiceResponseDTO struct {
	ID        string  `json:"id"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}
