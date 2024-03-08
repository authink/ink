package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type log interface {
	orm.Inserter[models.Log]
	orm.Finder[models.Log]
}

type logImpl app.AppContext

// Find implements log.
func (l *logImpl) Find(...model.Arg) (logs []models.Log, err error) {
	err = orm.Select(l.DB, sqls.Log.Find(), &logs, &model.Argument{})
	return
}

// Insert implements log.
func (l *logImpl) Insert(log *models.Log) error {
	return orm.NamedInsert(l.DB, sqls.Log.Insert(), log)
}

// InsertTx implements log.
func (l *logImpl) InsertTx(tx *sqlx.Tx, log *models.Log) error {
	return orm.NamedInsert(tx, sqls.Log.Insert(), log)
}

var _ log = (*logImpl)(nil)

func Log(appCtx *app.AppContext) log {
	return (*logImpl)(appCtx)
}
