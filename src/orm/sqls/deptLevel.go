package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type deptLevel interface {
	sql.Inserter
}

type deptLevelImpl struct{}

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.DeptLevel.Tname())).
		Columns(
			sbd.Field(db.DeptLevel.DeptId),
			sbd.Field(db.DeptLevel.SubDeptId),
		).
		String()
}

var DeptLevel deptLevel = &deptLevelImpl{}
