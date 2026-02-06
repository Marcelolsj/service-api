package dtos_users

type CreateUserRequestDTO struct {
	Nome        string `json:"nome"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
	TipoUsuario string `json:"tipo_usuario"`
}

type CreateUserResponseDTO struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	TipoUsuario string `json:"tipo_usuario"`
}
