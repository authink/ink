package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type deptLevel struct{}

// Delete implements inkstone.SQL.
func (d *deptLevel) Delete() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (d *deptLevel) Find() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (d *deptLevel) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (d *deptLevel) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, sub_dept_id) VALUES (:dept_id, :sub_dept_id)", table.DeptLevel)
}

// Save implements inkstone.SQL.
func (d *deptLevel) Save() string {
	panic("unimplemented")
}

// Update implements inkstone.SQL.
func (d *deptLevel) Update() string {
	panic("unimplemented")
}

var _ inkstone.SQL = (*deptLevel)(nil)
var DeptLevel = new(deptLevel)
