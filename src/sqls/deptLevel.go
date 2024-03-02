package sqls

import (
	"fmt"

	"github.com/authink/inkstone/sql"
)

type deptLevel interface {
	sql.Inserter
}

type deptLevelImpl struct{}

// Insert implements deptLevel.
func (d *deptLevelImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (dept_id, sub_dept_id) VALUES (:dept_id, :sub_dept_id)", table.DeptLevel)
}

var DeptLevel deptLevel = new(deptLevelImpl)
