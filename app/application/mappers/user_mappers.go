package mappers

import (
	dtos_users "service-api/application/dtos/users"
	"service-api/domain/models"
)

func MapCreateUserDTOToUser(createUserDto dtos_users.CreateUserRequestDTO) models.UserModel {
	return models.UserModel{
		Nome:        createUserDto.Nome,
		Login:       createUserDto.Login,
		Email:       createUserDto.Email,
		Senha:       createUserDto.Senha,
		TipoUsuario: createUserDto.TipoUsuario,
	}
}

func MapUserToUserResponseDTO(user models.UserModel) *dtos_users.UserResponseDTO {
	return &dtos_users.UserResponseDTO{
		ID:          user.ID,
		Nome:        user.Nome,
		Login:       user.Login,
		Email:       user.Email,
		TipoUsuario: user.TipoUsuario,
	}
}

func MapUsersToUserResponsesDTO(users []models.UserModel) []dtos_users.UserResponseDTO {
	var usersResponse []dtos_users.UserResponseDTO

	for _, user := range users {
		usersResponse = append(usersResponse, *MapUserToUserResponseDTO(user))
	}

	return usersResponse
}

func MapUpdateUserDTOToUser(updateUserDto dtos_users.UpdateUserRequestDTO, id string) models.UserModel {
	return models.UserModel{
		ID:    id,
		Nome:  updateUserDto.Nome,
		Login: updateUserDto.Login,
		Email: updateUserDto.Email,
	}
}
