package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/tables"
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
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s WHERE type = :type AND app_id = :app_id", tables.Group.TbName())
}

// Get implements group.
func (g *groupImpl) Get() string {
	return fmt.Sprintf("SELECT id, name, type, app_id, active FROM %s WHERE id = ?", tables.Group.TbName())
}

// GetForUpdate implements group.
func (g *groupImpl) GetForUpdate() string {
	return fmt.Sprintf("SELECT id, name, type, app_id, active FROM %s WHERE id = ? FOR UPDATE", tables.Group.TbName())
}

// Insert implements group.
func (g *groupImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (name, type, app_id) VALUES (:name, :type, :app_id)", tables.Group.TbName())
}

// Pagination implements group.
func (g *groupImpl) Pagination() string {
	return fmt.Sprintf("SELECT g.id, g.created_at, g.updated_at, g.name, g.type, g.app_id, a.name app_name, g.active FROM %s g, %s a WHERE g.app_id = a.id AND g.type = :type AND g.app_id = :app_id ORDER BY g.id DESC LIMIT :limit OFFSET :offset", tables.Group.TbName(), tables.TB_APPS)
}

// Save implements group.
func (g *groupImpl) Save() string {
	return fmt.Sprintf("INSERT INTO %s (name, type, app_id) VALUES (:name, :type, :app_id) ON DUPLICATE KEY UPDATE name = :name, active = :active", tables.Group.TbName())
}

// Update implements group.
func (g *groupImpl) Update() string {
	return fmt.Sprintf("UPDATE %s SET name = :name, active = :active WHERE id = :id", tables.Group.TbName())
}

var Group group = new(groupImpl)
