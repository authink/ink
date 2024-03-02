package sql

import (
	"fmt"

	"github.com/authink/inkstone/sql"
)

type deptStaff interface {
	sql.Inserter
}

type deptStaffImpl struct{}

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, staff_id) VALUES (:dept_id, :staff_id)", table.DeptStaff)
}

var DeptStaff deptStaff = new(deptStaffImpl)
