package sqls

import (
	"fmt"

	"github.com/authink/inkstone/sql"
)

type app interface {
	sql.Inserter
	sql.Saver
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Finder
}

type appImpl struct{}

// Find implements app.
func (a *appImpl) Find() string {
	return fmt.Sprintf("SELECT id, created_at, updated_at, name, active FROM %s ORDER BY id ASC", table.App)
}

// Get implements app.
func (a *appImpl) Get() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ?", table.App)
}

// GetForUpdate implements app.
func (a *appImpl) GetForUpdate() string {
	return fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ? FOR UPDATE", table.App)
}

// Insert implements app.
func (a *appImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", table.App)
}

// Save implements app.
func (a *appImpl) Save() string {
	return fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret) ON DUPLICATE KEY UPDATE active = :active, secret = :secret", table.App)
}

// Update implements app.
func (a *appImpl) Update() string {
	return fmt.Sprintf("UPDATE %s SET active = :active, secret = :secret WHERE id = :id", table.App)
}

var App app = new(appImpl)
