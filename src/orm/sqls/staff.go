package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
)

type staff interface {
	sql.Inserter
	sql.Saver
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Counter
	sql.Pager
	GetByEmail() string
}

type staffImpl struct{}

// Count implements staff.
func (s *staffImpl) Count() string {
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s", db.Staff.Tname())
}

// Get implements staff.
func (s *staffImpl) Get() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE id = ?", db.Staff.Tname())
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail() string {
	return fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE email = ?", db.Staff.Tname())
}

// GetForUpdate implements staff.
func (s *staffImpl) GetForUpdate() string {
	return fmt.Sprintf("SELECT id, email, phone, super, active, departure FROM %s WHERE id = ? FOR UPDATE", db.Staff.Tname())
}

// Insert implements staff.
func (s *staffImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", db.Staff.Tname())
}

// Pagination implements staff.
func (s *staffImpl) Pagination() string {
	return fmt.Sprintf("SELECT id, created_at, updated_at, email, phone, super, active, departure FROM %s ORDER BY id DESC LIMIT :limit OFFSET :offset", db.Staff.Tname())
}

// Save implements staff.
func (s *staffImpl) Save() string {
	return fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super) ON DUPLICATE KEY UPDATE password = :password, phone = :phone, super = :super, active = :active, departure = :departure", db.Staff.Tname())
}

// Update implements staff.
func (s *staffImpl) Update() string {
	return fmt.Sprintf("UPDATE %s SET password = :password, phone = :phone, super = :super, active = :active, departure = :departure WHERE id = :id", db.Staff.Tname())
}

var Staff staff = &staffImpl{}
