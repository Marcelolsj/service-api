package repository

import (
	"context"
	"service-api/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	UpdateUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*models.UserModel, error)
	GetUsers(ctx context.Context) ([]models.UserModel, error)
}
