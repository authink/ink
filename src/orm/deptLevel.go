package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type deptLevel interface {
	orm.Inserter[models.DeptLevel]
}

type deptLevelImpl app.AppContext

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert(deptLevel *models.DeptLevel) error {
	return namedExec(d.DB, sqls.DeptLevel.Insert(), deptLevel, handleInsertResult)
}

// InsertWithTx implements deptLevel.
func (d *deptLevelImpl) InsertTx(tx *sqlx.Tx, deptLevel *models.DeptLevel) error {
	return namedExec(tx, sqls.DeptLevel.Insert(), deptLevel, handleInsertResult)
}

var _ deptLevel = (*deptLevelImpl)(nil)

func DeptLevel(appCtx *app.AppContext) deptLevel {
	return (*deptLevelImpl)(appCtx)
}
