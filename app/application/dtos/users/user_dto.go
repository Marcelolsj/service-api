package dtos_users

type CreateUserRequestDTO struct {
	Nome        string `json:"nome"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
	TipoUsuario string `json:"tipo_usuario"`
}

type UserResponseDTO struct {
	ID          string `json:"id"`
	Nome        string `json:"nome"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	TipoUsuario string `json:"tipo_usuario"`
}

type UpdateUserRequestDTO struct {
	Nome  string `json:"nome"`
	Login string `json:"login"`
	Email string `json:"email"`
}
