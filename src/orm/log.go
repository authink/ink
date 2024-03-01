package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type log interface {
	inkstone.ORM[model.Log]
	GetWithTx(int, *sqlx.Tx) (*model.Log, error)
}

type logImpl inkstone.AppContext

// Delete implements log.
func (l *logImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements log.
func (l *logImpl) Find() (logs []model.Log, err error) {
	err = l.DB.Select(
		&logs,
		sql.Log.Find(),
	)
	return
}

// Get implements log.
// Subtle: this method shadows the method (*DB).Get of logImpl.DB.
func (l *logImpl) Get(int) (*model.Log, error) {
	panic("unimplemented")
}

// GetWithTx implements log.
func (l *logImpl) GetWithTx(int, *sqlx.Tx) (*model.Log, error) {
	panic("unimplemented")
}

// Insert implements log.
func (l *logImpl) Insert(log *model.Log) error {
	return namedExec(l.DB, sql.Log.Insert(), log, handleInsertResult)
}

// InsertWithTx implements log.
func (l *logImpl) InsertWithTx(*model.Log, *sqlx.Tx) error {
	panic("unimplemented")
}

// Save implements log.
func (l *logImpl) Save(*model.Log) error {
	panic("unimplemented")
}

// SaveWithTx implements log.
func (l *logImpl) SaveWithTx(*model.Log, *sqlx.Tx) error {
	panic("unimplemented")
}

// Update implements log.
func (l *logImpl) Update(*model.Log) error {
	panic("unimplemented")
}

// UpdateWithTx implements log.
func (l *logImpl) UpdateWithTx(*model.Log, *sqlx.Tx) error {
	panic("unimplemented")
}

var _ log = (*logImpl)(nil)

func Log(appCtx *inkstone.AppContext) log {
	return (*logImpl)(appCtx)
}
