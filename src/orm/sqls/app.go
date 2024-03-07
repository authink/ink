package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
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
func (a *appImpl) Find() (statement string) {
	statement, _ = sqlbuilder.
		Select(
			sql.Id,
			sql.CreatedAt,
			sql.UpdatedAt,
			db.App.Name,
			db.App.Active,
		).
		From(db.App.Tname()).
		OrderBy(sql.Id).
		Asc().
		Build()
	return statement
}

// Get implements app.
func (a *appImpl) Get() (statement string) {
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Id,
			db.App.Name,
			db.App.Secret,
			db.App.Active,
		).
		From(db.App.Tname()).
		Where(sb.EQ(sql.Id, sql.Named(sql.Id))).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// GetForUpdate implements app.
func (a *appImpl) GetForUpdate() (statement string) {
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Id,
			db.App.Name,
			db.App.Secret,
			db.App.Active,
		).
		From(db.App.Tname()).
		Where(sb.EQ(sql.Id, sql.Named(sql.Id))).
		ForUpdate().
		Build()
	return sql.ReplaceAtWithColon(statement)
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
