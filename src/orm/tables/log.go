package tables

import "github.com/authink/inkstone/orm/sql"

type log struct {
	Detail string
}

// TbName implements sql.Table.
func (a *log) TbName() string {
	return TB_LOGS
}

var _ sql.Table = (*log)(nil)
