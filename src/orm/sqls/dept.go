package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Dept struct {
	sql.SQLBase
}

func (d *Dept) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Department.Tname())).
		Columns(
			sbd.Field(db.Department.Name),
			sbd.Field(db.Department.OwnerId),
		).
		String()
}
