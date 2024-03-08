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
		InsertInto(sbd.Table(db.DeptLevel.Tname())).
		Columns(
			sbd.Field(db.DeptLevel.DeptId),
			sbd.Field(db.DeptLevel.SubDeptId),
		).
		String()
}
