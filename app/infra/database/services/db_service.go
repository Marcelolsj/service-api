package database_services

type DBService struct {
	ID        string  `db:"id"`
	Descricao string  `db:"descricao"`
	Valor     float64 `db:"valor"`
}
