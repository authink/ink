package tables

import "github.com/authink/inkstone/orm/sql"

type staff struct {
	Email     string
	Password  string
	Phone     string
	Super     string
	Active    string
	Departure string
}

// TbName implements sql.Table.
func (a *staff) TbName() string {
	return TB_STAFFS
}

var _ sql.Table = (*staff)(nil)
