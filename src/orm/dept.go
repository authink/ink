package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type dept interface {
	orm.Inserter[models.Department]
}

type deptImpl app.AppContext

// Insert implements dept.
func (d *deptImpl) Insert(dept *models.Department) error {
	return namedExec(d.DB, sqls.Dept.Insert(), dept, handleInsertResult)
}

// InsertWithTx implements dept.
func (d *deptImpl) InsertTx(tx *sqlx.Tx, dept *models.Department) error {
	return namedExec(tx, sqls.Dept.Insert(), dept, handleInsertResult)
}

var _ dept = (*deptImpl)(nil)

func Dept(appCtx *app.AppContext) dept {
	return (*deptImpl)(appCtx)
}
