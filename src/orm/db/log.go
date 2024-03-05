package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type log struct {
	Detail string
}

// Tname implements db.Table.
func (a *log) Tname() string {
	return "s_logs"
}

var _ db.Table = (*log)(nil)

var Log log

func init() {
	db.Bind(&Log, &models.Log{})
}
