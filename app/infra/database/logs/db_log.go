package database_logs

import "time"

type DBLog struct {
	ID        string    `db:"id"`
	IDRecurso string    `db:"id_recurso"`
	Log       string    `db:"log"`
	CriadoPor string    `db:"criado_por"`
	CriadoEm  time.Time `db:"criado_em"`
}
