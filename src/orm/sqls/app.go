package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	sbd "github.com/authink/sqlbuilder"
	"github.com/huandu/go-sqlbuilder"
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
	return sbd.
		NewBuilder().
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
func (a *appImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.App.Tname()).
		Cols(
			db.App.Name,
			db.App.Secret,
		).
		Values(
			sql.Named(db.App.Name),
			sql.Named(db.App.Secret),
		).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// Update implements app.
func (a *appImpl) Update() (statement string) {
	sb := sqlbuilder.NewUpdateBuilder()
	statement, _ = sb.
		Update(db.App.Tname()).
		Set(
			sb.Assign(
				db.App.Active,
				sql.Named(db.App.Active),
			),
			sb.Assign(
				db.App.Secret,
				sql.Named(db.App.Secret),
			),
		).
		Where(
			sb.EQ(sql.Id, sql.Named(sql.Id)),
		).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

var App app = &appImpl{}
