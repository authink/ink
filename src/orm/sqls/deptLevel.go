package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type DeptLevel struct {
	sql.SQLBase
}

func (d *DeptLevel) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.DeptLevel.Tname()).
		Columns(
			db.DeptLevel.DeptId,
			db.DeptLevel.SubDeptId,
		).
		String()
}
