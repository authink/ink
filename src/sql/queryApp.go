package sql

import "fmt"

type app struct{}

// Delete implements query.
func (*app) Delete() string {
	panic("unimplemented")
}

// Get implements query.
func (*app) Get() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ?", table.App)
}

// Insert implements query.
func (q *app) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", table.App)
}

var _ query = (*app)(nil)
var App = &app{}
