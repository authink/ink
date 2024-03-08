package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type dept interface {
	sql.Inserter
}

type deptImpl struct{}

// Insert implements dept.
func (d *deptImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Department.Tname())).
		Columns(
			sbd.Field(db.Department.Name),
			sbd.Field(db.Department.OwnerId),
		).
		String()
}

var Dept dept = &deptImpl{}
