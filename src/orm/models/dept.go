package models

import (
	"github.com/authink/orm/model"
)

// @model
// @db s_departments
type Department struct {
	model.Base
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

// @model
// @embed Department
type DepartmentWithOwner struct {
	Department
	OwnerEmail string `db:"owner_email"`
}

// @model
// @db s_dept_levels
type DeptLevel struct {
	model.Record
	DeptId    uint32 `db:"dept_id"`
	SubDeptId uint32 `db:"sub_dept_id"`
}

// @model
// @db s_dept_staffs
type DeptStaff struct {
	model.Record
	DeptId  uint32 `db:"dept_id"`
	StaffId uint32 `db:"staff_id"`
}

type CheckUnique struct {
	model.Argument
	Name string
}
