package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type deptStaff struct {
	DeptId  string
	StaffId string
}

// Tname implements db.Table.
func (d *deptStaff) Tname() string {
	return "s_dept_staffs"
}

var _ db.Table = (*deptStaff)(nil)

var DeptStaff deptStaff

func init() {
	db.Bind(&DeptStaff, &models.DeptStaff{})
}
