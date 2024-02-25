package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type app struct{}

func (*app) GetForUpdate() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ? FOR UPDATE", table.App)
}

// Update implements inkstone.SQL.
func (*app) Update() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (*app) Find() string {
	return fmt.Sprintf("SELECT id, created_at, updated_at, name, active FROM %s ORDER BY id ASC", table.App)
}

// Delete implements inkstone.SQL.
func (*app) Delete() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (*app) Get() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ?", table.App)
}

// Insert implements inkstone.SQL.
func (*app) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret) ON DUPLICATE KEY UPDATE active = :active, secret = :secret", table.App)
}

var _ inkstone.SQL = (*app)(nil)
var App = &app{}
