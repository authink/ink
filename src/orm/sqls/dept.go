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
		InsertInto(db.Department.Tname()).
		Columns(
			db.Department.Name,
			db.Department.OwnerId,
		).
		String()
}
