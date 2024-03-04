package tables

import "github.com/authink/inkstone/orm/sql"

type group struct {
	Name   string
	Type   string
	AppId  string
	Active string
}

// TbName implements sql.Table.
func (a *group) TbName() string {
	return TB_GOUPS
}

var _ sql.Table = (*group)(nil)
