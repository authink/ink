package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
)

type deptStaff interface {
	sql.Inserter
}

type deptStaffImpl struct{}

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, staff_id) VALUES (:dept_id, :staff_id)", db.DeptStaff.Tname())
}

var DeptStaff deptStaff = new(deptStaffImpl)
