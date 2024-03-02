package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type log interface {
	orm.Inserter[model.Log]
	orm.Finder[model.Log]
}

type logImpl a.AppContext

// Find implements log.
func (l *logImpl) Find(...any) (logs []model.Log, err error) {
	err = l.DB.Select(
		&logs,
		sql.Log.Find(),
	)
	return
}

// Insert implements log.
func (l *logImpl) Insert(log *model.Log) error {
	return namedExec(l.DB, sql.Log.Insert(), log, handleInsertResult)
}

// InsertTx implements log.
func (l *logImpl) InsertTx(tx *sqlx.Tx, log *model.Log) error {
	return namedExec(tx, sql.Log.Insert(), log, handleInsertResult)
}

var _ log = (*logImpl)(nil)

func Log(appCtx *a.AppContext) log {
	return (*logImpl)(appCtx)
}
