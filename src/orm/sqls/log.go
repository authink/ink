package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type log interface {
	sql.Inserter
	sql.Finder
}

type logImpl struct{}

// Find implements log.
func (l *logImpl) Find() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sbd.Field(db.Log.Detail),
		).
		From(sbd.Table(db.Log.Tname())).
		OrderBy(sql.Id).
		Desc().
		String()
}

// Insert implements log.
func (l *logImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Log.Tname())).
		Columns(sbd.Field(db.Log.Detail)).
		String()
}

var Log log = &logImpl{}
