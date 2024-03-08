package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type App struct {
	sql.SQLBase
}

func (a *App) Find() string {
	return sbd.NewBuilder().
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
		String()
}

func (a *App) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.App.Name,
			db.App.Secret,
			db.App.Active,
		).
		From(db.App.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (a *App) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.App.Name,
			db.App.Secret,
			db.App.Active,
		).
		From(db.App.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

func (a *App) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.App.Tname()).
		Columns(
			db.App.Name,
			db.App.Secret,
		).
		String()
}

func (a *App) Update() string {
	return sbd.NewBuilder().
		Update(db.App.Tname()).
		Set(
			db.App.Active,
			db.App.Secret,
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}
