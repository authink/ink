// Auto generated by inkstone, please do not change anything in this file
package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type deptLevel struct {
	DeptId string

	SubDeptId string
}

// Tname implements db.Table.
func (*deptLevel) Tname() string {
	return "s_dept_leves"
}

var _ db.Table = (*deptLevel)(nil)

var DeptLevel deptLevel

func init() {
	db.Bind(&DeptLevel, &models.DeptLevel{})
}
