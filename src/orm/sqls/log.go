package sqls

import (
	"github.com/authink/ink/src/orm/db"
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
			db.Log.Detail,
		).
		From(db.Log.Tname()).
		OrderBy(sql.Id).
		Desc().
		String()
}

func (l *Log) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.Log.Tname()).
		Columns(db.Log.Detail).
		String()
}
