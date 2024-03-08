package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/jmoiron/sqlx"
)

type deptStaff interface {
	orm.Inserter[models.DeptStaff]
}

type deptStaffImpl app.AppContext

// Insert implements deptStaff.
func (d *deptStaffImpl) Insert(deptStaff *models.DeptStaff) error {
	return orm.NamedInsert(d.DB, sqls.DeptStaff.Insert(), deptStaff)
}

// InsertWithTx implements deptStaff.
func (d *deptStaffImpl) InsertTx(tx *sqlx.Tx, deptStaff *models.DeptStaff) error {
	return orm.NamedInsert(tx, sqls.DeptStaff.Insert(), deptStaff)
}

var _ deptStaff = (*deptStaffImpl)(nil)

func DeptStaff(appCtx *app.AppContext) deptStaff {
	return (*deptStaffImpl)(appCtx)
}
