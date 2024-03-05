package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

type dept interface {
	sql.Inserter
}

type deptImpl struct{}

// Insert implements dept.
func (d *deptImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.InsertInto(db.Department.Tname()).Cols(db.Department.Name, db.Department.OwnerId).Values(
		sql.Named(db.Department.Name),
		sql.Named(db.Department.OwnerId),
	).Build()
	return sql.ReplaceAtWithColon(statement)
}

var Dept dept = &deptImpl{}
