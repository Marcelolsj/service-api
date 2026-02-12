package database_users

type DBUser struct {
	ID          string `db:"id"`
	Nome        string `db:"nome"`
	Login       string `db:"login"`
	Email       string `db:"email"`
	TipoUsuario string `db:"tipo_usuario"`
}
