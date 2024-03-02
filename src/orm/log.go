package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type log interface {
	orm.Inserter[models.Log]
	orm.Finder[models.Log]
}

type logImpl app.AppContext

// Find implements log.
func (l *logImpl) Find(args ...any) (logs []models.Log, err error) {
	err = doSelect(l.DB, &logs, sqls.Log.Find(), args...)
	return
}

// Insert implements log.
func (l *logImpl) Insert(log *models.Log) error {
	return namedExec(l.DB, sqls.Log.Insert(), log, afterInsert)
}

// InsertTx implements log.
func (l *logImpl) InsertTx(tx *sqlx.Tx, log *models.Log) error {
	return namedExec(tx, sqls.Log.Insert(), log, afterInsert)
}

var _ log = (*logImpl)(nil)

func Log(appCtx *app.AppContext) log {
	return (*logImpl)(appCtx)
}
