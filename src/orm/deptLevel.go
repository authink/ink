package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type deptLevel interface {
	orm.Inserter[model.DeptLevel]
}

type deptLevelImpl a.AppContext

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert(deptLevel *model.DeptLevel) error {
	return namedExec(d.DB, sql.DeptLevel.Insert(), deptLevel, handleInsertResult)
}

// InsertWithTx implements deptLevel.
func (d *deptLevelImpl) InsertTx(tx *sqlx.Tx, deptLevel *model.DeptLevel) error {
	return namedExec(tx, sql.DeptLevel.Insert(), deptLevel, handleInsertResult)
}

var _ deptLevel = (*deptLevelImpl)(nil)

func DeptLevel(appCtx *a.AppContext) deptLevel {
	return (*deptLevelImpl)(appCtx)
}
