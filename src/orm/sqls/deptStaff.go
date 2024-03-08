package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type deptStaff interface {
	sql.Inserter
}

type deptStaffImpl struct{}

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.DeptStaff.Tname())).
		Columns(
			sbd.Field(db.DeptStaff.DeptId),
			sbd.Field(db.DeptStaff.StaffId),
		).
		String()
}

var DeptStaff deptStaff = &deptStaffImpl{}
