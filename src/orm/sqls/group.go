package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type group interface {
	sql.Inserter
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Counter
	sql.Pager
}

type groupImpl struct{}

// Count implements group.
func (g *groupImpl) Count() string {
	return sbd.NewBuilder().
		Select(sbd.Field(sql.Id).Count()).
		From(sbd.Table(db.Group.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.Group.AppId)}).
		And(sbd.Equal{Left: sbd.Field(db.Group.Type)}).
		String()
}

// Get implements group.
func (g *groupImpl) Get() string {
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

// GetForUpdate implements group.
func (g *groupImpl) GetForUpdate() string {
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

// Insert implements group.
func (g *groupImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Group.Tname())).
		Columns(
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Type),
			sbd.Field(db.Group.AppId),
		).
		String()
}

// Pagination implements group.
func (g *groupImpl) Pagination() string {
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
			sbd.Field(db.App.Name).Of(aa),
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
		String()
}

// Update implements group.
func (g *groupImpl) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.Group.Tname())).
		Set(
			sbd.Field(db.Group.Name),
			sbd.Field(db.Group.Active),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

var Group group = &groupImpl{}
