package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type deptStaff interface {
	orm.Inserter[models.DeptStaff]
}

type deptStaffImpl app.AppContext

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert(deptStaff *models.DeptStaff) error {
	return namedExec(d.DB, sqls.DeptStaff.Insert(), deptStaff, afterInsert)
}

// InsertWithTx implements deptStaff.
func (d *deptStaffImpl) InsertTx(tx *sqlx.Tx, deptStaff *models.DeptStaff) error {
	return namedExec(tx, sqls.DeptStaff.Insert(), deptStaff, afterInsert)
}

var _ deptStaff = (*deptStaffImpl)(nil)

func DeptStaff(appCtx *app.AppContext) deptStaff {
	return (*deptStaffImpl)(appCtx)
}
