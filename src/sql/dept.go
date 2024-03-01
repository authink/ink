package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type dept struct{}

// Delete implements inkstone.SQL.
func (d *dept) Delete() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (d *dept) Find() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (d *dept) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (d *dept) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, owner_id) VALUES (:name, :owner_id)", table.Dept)
}

// Save implements inkstone.SQL.
func (d *dept) Save() string {
	panic("unimplemented")
}

// Update implements inkstone.SQL.
func (d *dept) Update() string {
	panic("unimplemented")
}

var _ inkstone.SQL = (*dept)(nil)
var Dept = new(dept)
