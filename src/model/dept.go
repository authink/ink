package model

import "github.com/authink/inkstone/orm"

type Department struct {
	orm.Model
	Name    string
	OwnerId uint32 `db:"owner_id"`
	Active  bool
}

func NewDept(name string, ownerId int) *Department {
	return &Department{
		Name:    name,
		OwnerId: uint32(ownerId),
		Active:  true,
	}
}

type DeptLevel struct {
	orm.Record
	DeptId    uint32 `db:"dept_id"`
	SubDeptId uint32 `db:"sub_dept_id"`
}

type DeptStaff struct {
	orm.Record
	DeptId  uint32 `db:"dept_id"`
	StaffId uint32 `db:"staff_id"`
}
