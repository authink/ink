package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type dept struct {
	Name    string
	OwnerId string
	Active  string
}

// Tname implements db.Table.
func (d *dept) Tname() string {
	return "s_departments"
}

var _ db.Table = (*dept)(nil)

var Dept dept

func init() {
	db.Bind(&Dept, &models.Department{})
}
