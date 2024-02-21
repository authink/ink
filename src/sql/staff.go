package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type staff struct{}

// Update implements inkstone.SQL.
func (*staff) Update() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (*staff) Find() string {
	panic("unimplemented")
}

// Delete implements inkstone.SQL.
func (*staff) Delete() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (*staff) Get() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE id = ?", table.Staff)
}

// Insert implements inkstone.SQL.
func (*staff) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", table.Staff)
}

func (*staff) GetByEmail() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE email = ?", table.Staff)
}

var _ inkstone.SQL = (*staff)(nil)
var Staff = &staff{}
