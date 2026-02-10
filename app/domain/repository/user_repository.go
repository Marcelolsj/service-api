package repository

import (
	"context"
	"service-api/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}
