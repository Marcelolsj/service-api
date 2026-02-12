package services

import (
	"context"
	dtos_users "service-api/application/dtos/users"
	http_errors "service-api/application/errors"
	"service-api/application/mappers"
	"service-api/domain/use_case"
)

type UserService struct {
	UserUseCase use_case.IUserUseCase
}

func (this *UserService) CreateUser(ctx context.Context, createUserDto dtos_users.CreateUserRequestDTO) (*dtos_users.UserResponseDTO, *http_errors.HttpError) {
	user := mappers.MapCreateUserDTOToUser(createUserDto)

	if err := this.UserUseCase.CreateUser(ctx, &user); err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapUserToUserResponseDTO(user), nil
}

func (this *UserService) UpdateUser(ctx context.Context, updateUserDto dtos_users.UpdateUserRequestDTO, id string) (*dtos_users.UserResponseDTO, *http_errors.HttpError) {
	user := mappers.MapUpdateUserDTOToUser(updateUserDto, id)

	if err := this.UserUseCase.UpdateUser(ctx, &user); err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapUserToUserResponseDTO(user), nil
}

func (this *UserService) DeleteUser(ctx context.Context, id string) *http_errors.HttpError {
	if err := this.UserUseCase.DeleteUser(ctx, id); err != nil {
		return http_errors.NewHttpError(err)
	}

	return nil
}

func (this *UserService) GetUsers(ctx context.Context) ([]dtos_users.UserResponseDTO, *http_errors.HttpError) {
	users, err := this.UserUseCase.GetUsers(ctx)
	if err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapUsersToUserResponsesDTO(users), nil
}

func (this *UserService) GetUser(ctx context.Context, id string) (*dtos_users.UserResponseDTO, *http_errors.HttpError) {
	user, err := this.UserUseCase.GetUser(ctx, id)
	if err != nil {
		return nil, http_errors.NewHttpError(err)
	}

	return mappers.MapUserToUserResponseDTO(*user), nil
}

func NewUserService(userUseCase use_case.IUserUseCase) *UserService {
	return &UserService{
		UserUseCase: userUseCase,
	}
}
