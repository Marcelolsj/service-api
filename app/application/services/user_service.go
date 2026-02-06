package services

import (
	dtos_users "service-api/application/dtos/users"
	"service-api/application/mappers"
	"service-api/domain/use_case"
)

type UserService struct {
	UserUseCase use_case.IUserUseCase
}

func (this *UserService) CreateUser(createUserDto dtos_users.CreateUserRequestDTO) (*dtos_users.CreateUserResponseDTO, error) {
	user := mappers.MapCreateUserDTOToUser(createUserDto)

	err := this.UserUseCase.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return mappers.MapUserToCreateUserResponseDTO(user), nil
}

func NewUserService(userUseCase use_case.IUserUseCase) *UserService {
	return &UserService{
		UserUseCase: userUseCase,
	}
}
