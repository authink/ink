package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type log struct{}

// Delete implements inkstone.SQL.
func (l *log) Delete() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (l *log) Find() string {
	panic("unimplemented")
}

// Get implements inkstone.SQL.
func (l *log) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (l *log) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (detail) VALUES (:detail)", table.Log)
}

// Save implements inkstone.SQL.
func (l *log) Save() string {
	panic("unimplemented")
}

// Update implements inkstone.SQL.
func (l *log) Update() string {
	panic("unimplemented")
}

var _ inkstone.SQL = (*log)(nil)
var Log = new(log)
