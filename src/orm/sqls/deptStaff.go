package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

type deptStaff interface {
	sql.Inserter
}

type deptStaffImpl struct{}

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.DeptStaff.Tname()).
		Cols(db.DeptStaff.DeptId, db.DeptStaff.StaffId).Values(db.DeptStaff.DeptId, db.DeptStaff.StaffId).Build()
	return sql.ReplaceAtWithColon(statement)
}

var DeptStaff deptStaff = &deptStaffImpl{}
