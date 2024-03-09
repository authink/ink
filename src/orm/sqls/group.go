package sqls

import (
	"github.com/authink/ink/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Group struct {
	sql.SQLBase
}

func (g *Group) Count() string {
	return sbd.NewBuilder().
		Select(sql.Id.Count()).
		From(db.Group.Tname()).
		Where(sbd.Equal{Left: db.Group.AppId}).
		And(sbd.Equal{Left: db.Group.Type}).
		String()
}

func (g *Group) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.Group.Name,
			db.Group.Type,
			db.Group.AppId,
			db.Group.Active,
		).
		From(db.Group.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (g *Group) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.Group.Name,
			db.Group.Type,
			db.Group.AppId,
			db.Group.Active,
		).
		From(db.Group.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

func (g *Group) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.Group.Tname()).
		Columns(
			db.Group.Name,
			db.Group.Type,
			db.Group.AppId,
		).
		String()
}

func (g *Group) Pagination() string {
	ag := sbd.Table("g")
	aa := sbd.Table("a")
	return sbd.NewBuilder().
		Select(
			sql.Id.Of(ag),
			sql.CreatedAt.Of(ag),
			sql.UpdatedAt.Of(ag),
			db.Group.Name.Of(ag),
			db.Group.Type.Of(ag),
			db.Group.AppId.Of(ag),
			db.App.Name.Of(aa).As(db.GroupWithApp.AppName),
			db.Group.Active.Of(ag),
		).
		From(
			db.Group.Tname().As(ag),
			db.App.Tname().As(aa),
		).
		Where(sbd.Equal{
			Left:  db.Group.AppId.Of(ag),
			Right: sql.Id.Of(aa),
		}).
		And(sbd.Equal{
			Left:  db.Group.AppId.Of(ag),
			Right: db.Group.AppId.Named(),
		}).
		And(sbd.Equal{
			Left:  db.Group.Type.Of(ag),
			Right: db.Group.Type.Named(),
		}).
		OrderBy(sql.Id.Of(ag)).
		Desc().
		Limit().
		String()
}

func (g *Group) Update() string {
	return sbd.NewBuilder().
		Update(db.Group.Tname()).
		Set(
			db.Group.Name,
			db.Group.Active,
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}
