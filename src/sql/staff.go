package sql

import "fmt"

type staff struct{}

// Delete implements SQL.
func (*staff) Delete() string {
	panic("unimplemented")
}

// Get implements SQL.
func (*staff) Get() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE id = ?", table.Staff)
}

// Insert implements SQL.
func (*staff) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", table.Staff)
}

func (*staff) GetByEmail() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE email = ?", table.Staff)
}

var _ SQL = (*staff)(nil)
var Staff = &staff{}