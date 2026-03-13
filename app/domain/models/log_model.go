package models

import (
	"time"
)

type LogModel struct {
	ID        string
	IDRecurso string
	Log       string
	CriadoPor string
	CriadoEm  time.Time
}
