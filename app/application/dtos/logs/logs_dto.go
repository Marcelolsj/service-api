package dtos_logs

type LogResponseDTO struct {
	ID        string `json:"id"`
	Log       string `json:"log"`
	CriadoPor string `json:"criado_por"`
	CriadoEm  string `json:"criado_em"`
}
