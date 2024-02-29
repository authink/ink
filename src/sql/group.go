package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type group struct{}

func (*group) Count() string {
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s WHERE type = ? AND app_id = ?", table.Group)
}

func (*group) Pagination() string {
	return fmt.Sprintf("SELECT g.id, g.created_at, g.updated_at, g.name, g.type, g.app_id, a.name app_name, g.active FROM %s g, %s a WHERE g.app_id = a.id AND g.type = ? AND g.app_id = ? ORDER BY g.id DESC LIMIT ? OFFSET ?", table.Group, table.App)
}

// Delete implements inkstone.SQL.
func (*group) Delete() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (*group) Find() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (*group) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (*group) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, type, app_id) VALUES (:name, :type, :app_id) ON DUPLICATE KEY UPDATE name = :name, active = :active", table.Group)
}

// Update implements inkstone.SQL.
func (*group) Update() string {
	panic("unimplemented")
}

var _ inkstone.SQL = (*group)(nil)
var Group = new(group)
