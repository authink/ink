package tables

import "github.com/authink/inkstone/orm/sql"

type dept struct {
	Name    string
	OwnerId string
	Active  string
}

// TbName implements sql.Table.
func (d *dept) TbName() string {
	return TB_DEPTS
}

var _ sql.Table = (*dept)(nil)

type deptLevel struct {
	DeptId    string
	SubDeptId string
}

// TbName implements sql.Table.
func (d *deptLevel) TbName() string {
	return TB_DEPT_LEVELS
}

var _ sql.Table = (*deptLevel)(nil)

type deptStaff struct {
	DeptId  string
	StaffId string
}

// TbName implements sql.Table.
func (d *deptStaff) TbName() string {
	return TB_DEPT_STAFFS
}

var _ sql.Table = (*deptStaff)(nil)
