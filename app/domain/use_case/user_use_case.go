package use_case

import (
	"context"
	"errors"
	"service-api/domain/models"
	"service-api/domain/repository"
)

type IUserUseCase interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func (this UserUseCase) CreateUser(ctx context.Context, user *models.User) error {
	err := this.validateUser(user)
	if err != nil {
		return err
	}

	err = this.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return err
}

func (u *UserUseCase) validateUser(user *models.User) error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório")
	}
	if user.Login == "" {
		return errors.New("O login é obrigatório")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório")
	}
	if user.Senha == "" {
		return errors.New("A senha é obrigatória")
	}
	if user.TipoUsuario == "" {
		return errors.New("O tipo de usuário é obrigatório")
	}

	return nil
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}
