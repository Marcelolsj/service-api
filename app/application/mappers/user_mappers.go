package mappers

import (
	dtos_users "service-api/application/dtos/users"
	"service-api/domain/models"
)

func MapCreateUserDTOToUser(createUserDto dtos_users.CreateUserRequestDTO) *models.User {
	return &models.User{
		Nome:        createUserDto.Nome,
		Login:       createUserDto.Login,
		Email:       createUserDto.Email,
		Senha:       createUserDto.Senha,
		TipoUsuario: createUserDto.TipoUsuario,
	}
}

func MapUserToCreateUserResponseDTO(user *models.User) *dtos_users.CreateUserResponseDTO {
	return &dtos_users.CreateUserResponseDTO{
		ID:          user.ID,
		Nome:        user.Nome,
		Login:       user.Login,
		Email:       user.Email,
		TipoUsuario: user.TipoUsuario,
	}
}
