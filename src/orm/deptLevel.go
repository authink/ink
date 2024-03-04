package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
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
	return orm.NamedInsert(d.DB, sqls.DeptLevel.Insert(), deptLevel)
}

// InsertWithTx implements deptLevel.
func (d *deptLevelImpl) InsertTx(tx *sqlx.Tx, deptLevel *models.DeptLevel) error {
	return orm.NamedInsert(tx, sqls.DeptLevel.Insert(), deptLevel)
}

var _ deptLevel = (*deptLevelImpl)(nil)

func DeptLevel(appCtx *app.AppContext) deptLevel {
	return (*deptLevelImpl)(appCtx)
}
