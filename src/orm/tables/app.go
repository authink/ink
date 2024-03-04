package tables

import (
	"github.com/authink/inkstone/orm/sql"
)

type app struct {
	Name   string
	Secret string
	Active string
}

// TbName implements sql.Table.
func (a *app) TbName() string {
	return TB_APPS
}

var _ sql.Table = (*app)(nil)
