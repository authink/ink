package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
)

type group interface {
	sql.Inserter
	sql.Saver
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Counter
	sql.Pager
}

type groupImpl struct{}

// Count implements group.
func (g *groupImpl) Count() string {
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s WHERE type = :type AND app_id = :app_id", db.Group.Tname())
}

// Get implements group.
func (g *groupImpl) Get() string {
	return fmt.Sprintf("SELECT id, name, type, app_id, active FROM %s WHERE id = ?", db.Group.Tname())
}

// GetForUpdate implements group.
func (g *groupImpl) GetForUpdate() string {
	return fmt.Sprintf("SELECT id, name, type, app_id, active FROM %s WHERE id = ? FOR UPDATE", db.Group.Tname())
}

// Insert implements group.
func (g *groupImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, type, app_id) VALUES (:name, :type, :app_id)", db.Group.Tname())
}

// Pagination implements group.
func (g *groupImpl) Pagination() string {
	return fmt.Sprintf("SELECT g.id, g.created_at, g.updated_at, g.name, g.type, g.app_id, a.name app_name, g.active FROM %s g, %s a WHERE g.app_id = a.id AND g.type = :type AND g.app_id = :app_id ORDER BY g.id DESC LIMIT :limit OFFSET :offset", db.Group.Tname(), db.App.Tname())
}

// Save implements group.
func (g *groupImpl) Save() string {
	return fmt.Sprintf("INSERT INTO %s (name, type, app_id) VALUES (:name, :type, :app_id) ON DUPLICATE KEY UPDATE name = :name, active = :active", db.Group.Tname())
}

// Update implements group.
func (g *groupImpl) Update() string {
	return fmt.Sprintf("UPDATE %s SET name = :name, active = :active WHERE id = :id", db.Group.Tname())
}

var Group group = &groupImpl{}
