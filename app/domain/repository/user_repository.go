package repository

import "service-api/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
}
