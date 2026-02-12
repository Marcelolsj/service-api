package database_users

import (
	"context"
	"database/sql"
	"service-api/domain/models"
	"service-api/domain/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DbUserRepository struct {
	db *sqlx.DB
}

func (this *DbUserRepository) CreateUser(ctx context.Context, user *models.UserModel) error {
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

func (this *DbUserRepository) UpdateUser(ctx context.Context, user *models.UserModel) error {
	_, err := this.db.ExecContext(ctx,
		`update usuarios
		 set nome = ?, 
		 	 login = ?, 
			 email = ?
		 where id = ?`,
		user.Nome,
		user.Login,
		user.Email,
		user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *DbUserRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := this.db.ExecContext(ctx,
		`delete from usuarios where id = ?`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *DbUserRepository) GetUser(ctx context.Context, id string) (*models.UserModel, error) {
	var user DBUser

	query := `
		SELECT id, nome, login, email, tipo_usuario
		FROM usuarios
		WHERE id = ?
	`

	err := this.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	userModel := MapDBUserToUser(user)
	return &userModel, nil
}

func (this *DbUserRepository) GetUsers(ctx context.Context) ([]models.UserModel, error) {
	var users []DBUser

	query := `
		SELECT id, nome, login, email, tipo_usuario
		FROM usuarios
	`

	err := this.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	userModels := MapDBUsersToUsers(users)
	return userModels, nil
}

func NewDbUserRepository(db *sqlx.DB) repository.UserRepository {
	return &DbUserRepository{
		db: db,
	}
}
