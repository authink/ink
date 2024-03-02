package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type dept interface {
	orm.Inserter[model.Department]
}

type deptImpl a.AppContext

// Insert implements dept.
func (d *deptImpl) Insert(dept *model.Department) error {
	return namedExec(d.DB, sql.Dept.Insert(), dept, handleInsertResult)
}

// InsertWithTx implements dept.
func (d *deptImpl) InsertTx(tx *sqlx.Tx, dept *model.Department) error {
	return namedExec(tx, sql.Dept.Insert(), dept, handleInsertResult)
}

var _ dept = (*deptImpl)(nil)

func Dept(appCtx *a.AppContext) dept {
	return (*deptImpl)(appCtx)
}
