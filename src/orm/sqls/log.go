package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

type log interface {
	sql.Inserter
	sql.Finder
}

type logImpl struct{}

// Find implements log.
func (l *logImpl) Find() (statement string) {
	statement, _ = sqlbuilder.
		Select(
			sql.Id,
			sql.CreatedAt,
			db.Log.Detail,
		).
		From(db.Log.Tname()).
		OrderBy(sql.Id).
		Desc().
		Build()
	return statement
}

// Insert implements log.
func (l *logImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.Log.Tname()).
		Cols(db.Log.Detail).
		Values(sql.Named(db.Log.Detail)).
		Build()

	return sql.ReplaceAtWithColon(statement)
}

var Log log = &logImpl{}
