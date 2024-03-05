package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

type deptLevel interface {
	sql.Inserter
}

type deptLevelImpl struct{}

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.DeptLevel.Tname()).
		Cols(db.DeptLevel.DeptId, db.DeptLevel.SubDeptId).Values(
		sql.Named(db.DeptLevel.DeptId),
		sql.Named(db.DeptLevel.SubDeptId),
	).Build()
	return sql.ReplaceAtWithColon(statement)
}

var DeptLevel deptLevel = &deptLevelImpl{}
