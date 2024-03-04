package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/tables"
	"github.com/authink/inkstone/orm/sql"
)

type deptLevel interface {
	sql.Inserter
}

type deptLevelImpl struct{}

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, sub_dept_id) VALUES (:dept_id, :sub_dept_id)", tables.DeptLevel.TbName())
}

var DeptLevel deptLevel = new(deptLevelImpl)
