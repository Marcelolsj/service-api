package database_users

import (
	"context"
	"service-api/domain/models"
	"service-api/domain/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DbUserRepository struct {
	db *sqlx.DB
}

func (this *DbUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	id := uuid.New().String()

	_, err := this.db.ExecContext(ctx,
		`INSERT INTO usuarios (id, nome, login, email, senha, tipo_usuario)
		VALUES(?, ?, ?, ?, ?, ?) `,
		id,
		user.Nome,
		user.Login,
		user.Email,
		user.Senha,
		user.TipoUsuario,
	)
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func NewDbUserRepository(db *sqlx.DB) repository.UserRepository {
	return &DbUserRepository{
		db: db,
	}
}
