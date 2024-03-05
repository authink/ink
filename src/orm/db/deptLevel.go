package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type deptLevel struct {
	DeptId    string
	SubDeptId string
}

// Tname implements db.Table.
func (d *deptLevel) Tname() string {
	return "s_dept_levels"
}

var _ db.Table = (*deptLevel)(nil)

var DeptLevel deptLevel

func init() {
	db.Bind(&DeptLevel, &models.DeptLevel{})
}
