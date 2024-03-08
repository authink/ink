package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Group struct {
	sql.SQLBase
}

func (g *Group) Count() string {
	return sbd.NewBuilder().
		Select(sbd.Field(sql.Id).Count()).
		From(sbd.Table(db.Group.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.Group.AppId)}).
		And(sbd.Equal{Left: sbd.Field(db.Group.Type)}).
		String()
}

func (g *Group) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Type),
			sbd.Field(db.Group.AppId),
			sbd.Field(db.Group.Active),
		).
		From(sbd.Table(db.Group.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (g *Group) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Type),
			sbd.Field(db.Group.AppId),
			sbd.Field(db.Group.Active),
		).
		From(sbd.Table(db.Group.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

func (g *Group) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Group.Tname())).
		Columns(
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Type),
			sbd.Field(db.Group.AppId),
		).
		String()
}

func (g *Group) Pagination() string {
	ag := "g"
	aa := "a"
	fId := sbd.Field(sql.Id)
	fAppId := sbd.Field(db.Group.AppId)
	fType := sbd.Field(db.Group.Type)
	return sbd.NewBuilder().
		Select(
			fId.Of(ag),
			sbd.Field(sql.CreatedAt).Of(ag),
			sbd.Field(sql.UpdatedAt).Of(ag),
			sbd.Field(db.Group.Name).Of(ag),
			fType.Of(ag),
			fAppId.Of(ag),
			sbd.Field(db.App.Name).Of(aa).As(db.GroupWithApp.AppName),
			sbd.Field(db.Group.Active).Of(ag),
		).
		From(
			sbd.Table(db.Group.Tname()).As(ag),
			sbd.Table(db.App.Tname()).As(aa),
		).
		Where(sbd.Equal{
			Left:  fAppId.Of(ag),
			Right: fId.Of(aa),
		}).
		And(sbd.Equal{
			Left:  fAppId.Of(ag),
			Right: fAppId.Named(),
		}).
		And(sbd.Equal{
			Left:  fType.Of(ag),
			Right: fType.Named(),
		}).
		OrderBy(fId.Of(ag)).
		Desc().
		Limit().
		String()
}

func (g *Group) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.Group.Tname())).
		Set(
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Active),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}
