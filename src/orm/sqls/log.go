package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Log struct {
	sql.SQLBase
}

func (l *Log) Find() string {
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

func (l *Log) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Log.Tname())).
		Columns(sbd.Field(db.Log.Detail)).
		String()
}
