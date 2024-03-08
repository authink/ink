package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type DeptStaff struct {
	sql.SQLBase
}

func (d *DeptStaff) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.DeptStaff.Tname()).
		Columns(
			db.DeptStaff.DeptId,
			db.DeptStaff.StaffId,
		).
		String()
}
