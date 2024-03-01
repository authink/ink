package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type dept interface {
	inkstone.ORM[model.Department]
}

type deptImpl inkstone.AppContext

// Delete implements dept.
func (d *deptImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements dept.
func (d *deptImpl) Find() ([]model.Department, error) {
	panic("unimplemented")
}

// Get implements dept.
// Subtle: this method shadows the method (*DB).Get of deptImpl.DB.
func (d *deptImpl) Get(int) (*model.Department, error) {
	panic("unimplemented")
}

// Insert implements dept.
func (d *deptImpl) Insert(dept *model.Department) error {
	return namedExec(d.DB, sql.Dept.Insert(), dept, handleInsertResult)
}

// InsertWithTx implements dept.
func (d *deptImpl) InsertWithTx(*model.Department, *sqlx.Tx) error {
	panic("unimplemented")
}

// Save implements dept.
func (d *deptImpl) Save(*model.Department) error {
	panic("unimplemented")
}

// SaveWithTx implements dept.
func (d *deptImpl) SaveWithTx(*model.Department, *sqlx.Tx) error {
	panic("unimplemented")
}

// Update implements dept.
func (d *deptImpl) Update(*model.Department) error {
	panic("unimplemented")
}

// UpdateWithTx implements dept.
func (d *deptImpl) UpdateWithTx(*model.Department, *sqlx.Tx) error {
	panic("unimplemented")
}

var _ dept = (*deptImpl)(nil)

func Dept(appCtx *inkstone.AppContext) dept {
	return (*deptImpl)(appCtx)
}
