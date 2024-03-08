package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type staff interface {
	sql.Inserter
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
	return sbd.NewBuilder().
		Select(sbd.Field(sql.Id).Count()).
		From(sbd.Table(db.Staff.Tname())).
		String()
}

// Get implements staff.
func (s *staffImpl) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.Staff.Email)}).
		String()
}

// GetForUpdate implements staff.
func (s *staffImpl) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

// Insert implements staff.
func (s *staffImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Staff.Tname())).
		Columns(
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
		).
		String()
}

// Pagination implements staff.
func (s *staffImpl) Pagination() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sql.UpdatedAt,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
		).
		From(sbd.Table(db.Staff.Tname())).
		OrderBy(sql.Id).
		Desc().
		Limit().
		String()
}

// Update implements staff.
func (s *staffImpl) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.Staff.Tname())).
		Set(
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

var Staff staff = &staffImpl{}
