package sql

import (
	"fmt"

	"github.com/authink/inkstone/sql"
)

type dept interface {
	sql.Inserter
}

type deptImpl struct{}

// Insert implements dept.
func (d *deptImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, owner_id) VALUES (:name, :owner_id)", table.Dept)
}

var Dept dept = new(deptImpl)
