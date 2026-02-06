package database_users

import (
	"fmt"
	"service-api/domain/models"
	"service-api/domain/repository"
)

type DbUserRepository struct {
	// Database connection or ORM instance can be added here
}

func (db *DbUserRepository) CreateUser(user *models.User) error {
	fmt.Println("User created:", user)
	return nil
}

func NewDbUserRepository() repository.UserRepository {
	return &DbUserRepository{}
}
