package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type deptStaff struct{}

// Delete implements inkstone.SQL.
func (d *deptStaff) Delete() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (d *deptStaff) Find() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (d *deptStaff) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (d *deptStaff) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, staff_id) VALUES (:dept_id, :staff_id)", table.DeptStaff)
}

// Save implements inkstone.SQL.
func (d *deptStaff) Save() string {
	panic("unimplemented")
}

// Update implements inkstone.SQL.
func (d *deptStaff) Update() string {
	panic("unimplemented")
}

var _ inkstone.SQL = (*deptStaff)(nil)
var DeptStaff = new(deptStaff)
