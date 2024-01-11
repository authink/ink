package sql

import "fmt"

type staff struct{}

// Get implements query.
func (*staff) Get() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE email = ?", table.Staff)
}

// Insert implements query.
func (*staff) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", table.Staff)
}

var _ query = (*staff)(nil)
var Staff = &staff{}
