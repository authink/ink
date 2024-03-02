package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type deptStaff interface {
	orm.Inserter[model.DeptStaff]
}

type deptStaffImpl a.AppContext

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert(deptStaff *model.DeptStaff) error {
	return namedExec(d.DB, sql.DeptStaff.Insert(), deptStaff, handleInsertResult)
}

// InsertWithTx implements deptStaff.
func (d *deptStaffImpl) InsertTx(tx *sqlx.Tx, deptStaff *model.DeptStaff) error {
	return namedExec(tx, sql.DeptStaff.Insert(), deptStaff, handleInsertResult)
}

var _ deptStaff = (*deptStaffImpl)(nil)

func DeptStaff(appCtx *a.AppContext) deptStaff {
	return (*deptStaffImpl)(appCtx)
}
