package use_case

import (
	"context"
	domain_errors "service-api/domain/bu_errors"
	"service-api/domain/models"
	"service-api/domain/repository"
)

type IUserUseCase interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	UpdateUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*models.UserModel, error)
	GetUsers(ctx context.Context) ([]models.UserModel, error)
}

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func (this UserUseCase) CreateUser(ctx context.Context, user *models.UserModel) error {
	err := this.validateUser(user, true)
	if err != nil {
		return err
	}

	err = this.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	updatedUser, err := this.GetUser(ctx, user.ID)
	if updatedUser != nil {
		*user = *updatedUser
	} else {
		return domain_errors.NewBusinessErrorMsg("Usuário não encontrado após atualização")
	}

	return err
}

func (this UserUseCase) UpdateUser(ctx context.Context, user *models.UserModel) error {
	err := this.validateUser(user, false)
	if err != nil {
		return err
	}

	if user.ID == "" {
		return domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	err = this.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	updatedUser, err := this.GetUser(ctx, user.ID)
	if updatedUser != nil {
		*user = *updatedUser
	} else {
		return domain_errors.NewBusinessErrorMsg("Usuário não encontrado após atualização")
	}

	return err
}

func (this UserUseCase) DeleteUser(ctx context.Context, id string) error {
	err := this.UserRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	if id == "" {
		return domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	return err
}

func (this UserUseCase) GetUser(ctx context.Context, id string) (*models.UserModel, error) {
	user, err := this.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if id == "" {
		return nil, domain_errors.NewBusinessErrorMsg("O ID é obrigatório")
	}

	return user, nil
}

func (this UserUseCase) GetUsers(ctx context.Context) ([]models.UserModel, error) {
	users, err := this.UserRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCase) validateUser(user *models.UserModel, insert bool) error {
	if user.Nome == "" {
		return domain_errors.NewBusinessErrorMsg("O nome é obrigatório")
	}
	if user.Login == "" {
		return domain_errors.NewBusinessErrorMsg("O login é obrigatório")
	}
	if user.Email == "" {
		return domain_errors.NewBusinessErrorMsg("O email é obrigatório")
	}
	if user.Senha == "" && insert {
		return domain_errors.NewBusinessErrorMsg("A senha é obrigatória")
	}
	if user.TipoUsuario == "" && insert {
		return domain_errors.NewBusinessErrorMsg("O tipo de usuário é obrigatório")
	}

	return nil
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}
