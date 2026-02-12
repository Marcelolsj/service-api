package database_users

import "service-api/domain/models"

func MapDBUserToUser(dbUser DBUser) models.UserModel {
	return models.UserModel{
		ID:          dbUser.ID,
		Nome:        dbUser.Nome,
		Login:       dbUser.Login,
		Email:       dbUser.Email,
		TipoUsuario: dbUser.TipoUsuario,
	}
}

func MapDBUsersToUsers(dbUsers []DBUser) []models.UserModel {
	var users []models.UserModel

	for _, user := range dbUsers {
		users = append(users, MapDBUserToUser(user))
	}

	return users
}
