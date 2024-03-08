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
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Active),
		).
		From(sbd.Table(db.App.Tname())).
		OrderBy(sql.Id).
		Asc().
		String()
}

func (a *App) Get() string {
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

func (a *App) GetForUpdate() string {
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

func (a *App) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.App.Tname())).
		Columns(
			sbd.Field(db.App.Name),
			sbd.Field(db.App.Secret),
		).
		String()
}

func (a *App) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.App.Tname())).
		Set(
			sbd.Field(db.App.Active),
			sbd.Field(db.App.Secret),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}
