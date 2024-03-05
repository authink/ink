package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type staff struct {
	Email     string
	Password  string
	Phone     string
	Super     string
	Active    string
	Departure string
}

// Tname implements db.Table.
func (a *staff) Tname() string {
	return "s_staffs"
}

var _ db.Table = (*staff)(nil)

var Staff staff

func init() {
	db.Bind(&Staff, &models.Staff{})
}
