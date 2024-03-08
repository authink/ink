package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type app interface {
	sql.Inserter
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Finder
}

type appImpl struct{}

// Find implements app.
func (a *appImpl) Find() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sql.UpdatedAt,
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Active),
		).
		From(sbd.Table(db.App.Tname())).
		OrderBy(sql.Id).
		Asc().
		String()
}

// Get implements app.
func (a *appImpl) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Secret),
			sbd.Field(db.App.Active),
		).
		From(sbd.Table(db.App.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

// GetForUpdate implements app.
func (a *appImpl) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Secret),
			sbd.Field(db.App.Active),
		).
		From(sbd.Table(db.App.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

// Insert implements app.
func (a *appImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.App.Tname())).
		Columns(
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Secret),
		).
		String()
}

// Update implements app.
func (a *appImpl) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.App.Tname())).
		Set(
			sbd.Field(db.App.Active),
			sbd.Field(db.App.Secret),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

var App app = &appImpl{}
