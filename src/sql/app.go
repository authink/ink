package sql

import "fmt"

type app struct{}

// Find implements SQL.
func (*app) Find() string {
	return fmt.Sprintf("SELECT id, created_at, updated_at, name, active FROM %s ORDER BY id ASC", table.App)
}

// Delete implements SQL.
func (*app) Delete() string {
	panic("unimplemented")
}

// Get implements SQL.
func (*app) Get() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ?", table.App)
}

// Insert implements SQL.
func (*app) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", table.App)
}

var _ SQL = (*app)(nil)
var App = &app{}
