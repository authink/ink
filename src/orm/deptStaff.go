package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type deptStaff interface {
	inkstone.ORM[model.DeptStaff]
}

type deptStaffImpl inkstone.AppContext

// Delete implements deptStaff.
func (d *deptStaffImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements deptStaff.
func (d *deptStaffImpl) Find() ([]model.DeptStaff, error) {
	panic("unimplemented")
}

// Get implements deptStaff.
// Subtle: this method shadows the method (*DB).Get of deptStaffImpl.DB.
func (d *deptStaffImpl) Get(int) (*model.DeptStaff, error) {
	panic("unimplemented")
}

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert(deptStaff *model.DeptStaff) error {
	return namedExec(d.DB, sql.DeptStaff.Insert(), deptStaff, handleInsertResult)
}

// InsertWithTx implements deptStaff.
func (d *deptStaffImpl) InsertWithTx(*model.DeptStaff, *sqlx.Tx) error {
	panic("unimplemented")
}

// Save implements deptStaff.
func (d *deptStaffImpl) Save(*model.DeptStaff) error {
	panic("unimplemented")
}

// SaveWithTx implements deptStaff.
func (d *deptStaffImpl) SaveWithTx(*model.DeptStaff, *sqlx.Tx) error {
	panic("unimplemented")
}

// Update implements deptStaff.
func (d *deptStaffImpl) Update(*model.DeptStaff) error {
	panic("unimplemented")
}

// UpdateWithTx implements deptStaff.
func (d *deptStaffImpl) UpdateWithTx(*model.DeptStaff, *sqlx.Tx) error {
	panic("unimplemented")
}

var _ deptStaff = (*deptStaffImpl)(nil)

func DeptStaff(appCtx *inkstone.AppContext) deptStaff {
	return (*deptStaffImpl)(appCtx)
}
