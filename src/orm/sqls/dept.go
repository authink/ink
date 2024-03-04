package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/tables"
	"github.com/authink/inkstone/orm/sql"
)

type dept interface {
	sql.Inserter
}

type deptImpl struct{}

// Insert implements dept.
func (d *deptImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, owner_id) VALUES (:name, :owner_id)", tables.Dept.TbName())
}

var Dept dept = new(deptImpl)
