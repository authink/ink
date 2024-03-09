// Auto generated by authink/orm, please do not change anything in this file
package db

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/orm/db"
	sbd "github.com/authink/sqlbuilder"
)

type staff struct {
	Email sbd.Field

	Password sbd.Field

	Phone sbd.Field

	Super sbd.Field

	Active sbd.Field

	Departure sbd.Field
}

// Tname implements db.Table.
func (*staff) Tname() sbd.Table {
	return "s_staffs"
}

var _ db.Table = (*staff)(nil)

var Staff staff

func init() {
	db.Bind(&Staff, &models.Staff{})
}
