package sqls

import (
	"github.com/authink/ink.go/src/orm/tables"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

var (
	tbApp  = tables.App
	tbnApp = tbApp.TbName()
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
			tbApp.Name,
			tbApp.Active,
		).
		From(tbnApp).
		OrderBy(sql.Id).
		Asc().
		Build()
	return statement
}

// Get implements app.
func (a *appImpl) Get() (statement string) {
	statement, _ = sqlbuilder.
		Select(
			sql.Id,
			tbApp.Name,
			tbApp.Secret,
			tbApp.Active,
		).
		From(tbnApp).
		Where(sql.EQ(sql.Id, "?")).
		Build()
	return statement
}

// GetForUpdate implements app.
func (a *appImpl) GetForUpdate() (statement string) {
	statement, _ = sqlbuilder.
		Select(
			sql.Id,
			tbApp.Name,
			tbApp.Secret,
			tbApp.Active,
		).
		From(tbnApp).
		Where(sql.EQ(sql.Id, "?")).
		ForUpdate().
		Build()
	return statement
}

// Insert implements app.
func (a *appImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(tbnApp).
		Cols(
			tbApp.Name,
			tbApp.Secret,
		).
		Values(
			sql.Named(tbApp.Name),
			sql.Named(tbApp.Secret),
		).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// Update implements app.
func (a *appImpl) Update() (statement string) {
	sb := sqlbuilder.NewUpdateBuilder()
	statement, _ = sb.
		Update(tbnApp).
		Set(
			sb.Assign(
				tbApp.Active,
				sql.Named(tbApp.Active),
			),
			sb.Assign(
				tbApp.Secret,
				sql.Named(tbApp.Secret),
			),
		).
		Where(
			sb.EQ(sql.Id, sql.Named(sql.Id)),
		).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

var App app = new(appImpl)
