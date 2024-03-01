package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type deptLevel interface {
	inkstone.ORM[model.DeptLevel]
}

type deptLevelImpl inkstone.AppContext

// Delete implements deptLevel.
func (d *deptLevelImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements deptLevel.
func (d *deptLevelImpl) Find() ([]model.DeptLevel, error) {
	panic("unimplemented")
}

// Get implements deptLevel.
// Subtle: this method shadows the method (*DB).Get of deptLevelImpl.DB.
func (d *deptLevelImpl) Get(int) (*model.DeptLevel, error) {
	panic("unimplemented")
}

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert(deptLevel *model.DeptLevel) error {
	return namedExec(d.DB, sql.DeptLevel.Insert(), deptLevel, handleInsertResult)
}

// InsertWithTx implements deptLevel.
func (d *deptLevelImpl) InsertWithTx(*model.DeptLevel, *sqlx.Tx) error {
	panic("unimplemented")
}

// Save implements deptLevel.
func (d *deptLevelImpl) Save(*model.DeptLevel) error {
	panic("unimplemented")
}

// SaveWithTx implements deptLevel.
func (d *deptLevelImpl) SaveWithTx(*model.DeptLevel, *sqlx.Tx) error {
	panic("unimplemented")
}

// Update implements deptLevel.
func (d *deptLevelImpl) Update(*model.DeptLevel) error {
	panic("unimplemented")
}

// UpdateWithTx implements deptLevel.
func (d *deptLevelImpl) UpdateWithTx(*model.DeptLevel, *sqlx.Tx) error {
	panic("unimplemented")
}

var _ deptLevel = (*deptLevelImpl)(nil)

func DeptLevel(appCtx *inkstone.AppContext) deptLevel {
	return (*deptLevelImpl)(appCtx)
}
