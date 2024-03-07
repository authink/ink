package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
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
func (g *groupImpl) Count() (statement string) {
	tbnAlias := "g"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Count(tbnAlias),
		).
		From(sb.As(db.Group.Tname(), tbnAlias)).
		Where(
			sb.EQ(db.Group.Type, sql.Named(db.Group.Type)),
			sb.EQ(db.Group.AppId, sql.Named(db.Group.AppId)),
		).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// Get implements group.
func (g *groupImpl) Get() (statement string) {
	statement, _ = sqlbuilder.Select(
		sql.Id,
		db.Group.Name,
		db.Group.Type,
		db.Group.AppId,
		db.Group.Active,
	).From(db.Group.Tname()).Where(sql.EQ(sql.Id, "?")).Build()
	return statement
}

// GetForUpdate implements group.
func (g *groupImpl) GetForUpdate() (statement string) {
	statement, _ = sqlbuilder.Select(
		sql.Id,
		db.Group.Name,
		db.Group.Type,
		db.Group.AppId,
		db.Group.Active,
	).
		From(db.Group.Tname()).
		Where(sql.EQ(sql.Id, "?")).
		ForUpdate().
		Build()
	return statement
}

// Insert implements group.
func (g *groupImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.InsertInto(db.Group.Tname()).
		Cols(
			db.Group.Name,
			db.Group.Type,
			db.Group.AppId,
		).Values(
		sql.Named(db.Group.Name),
		sql.Named(db.Group.Type),
		sql.Named(db.Group.AppId),
	).Build()
	return sql.ReplaceAtWithColon(statement)
}

// Pagination implements group.
func (g *groupImpl) Pagination() (statement string) {
	tbnAlias1 := "g"
	tbnAlias2 := "a"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.Select(
		sql.Col(tbnAlias1, sql.Id),
		sql.Col(tbnAlias1, sql.CreatedAt),
		sql.Col(tbnAlias1, sql.UpdatedAt),
		sql.Col(tbnAlias1, db.Group.Name),
		sql.Col(tbnAlias1, db.Group.Type),
		sql.Col(tbnAlias1, db.Group.AppId),
		sql.Col(tbnAlias2, db.App.Name),
		sql.Col(tbnAlias1, db.Group.Active),
	).From(
		sb.As(db.Group.Tname(), tbnAlias1),
		sb.As(db.App.Tname(), tbnAlias2),
	).Where(
		sql.EQ(
			sql.Col(tbnAlias1, db.Group.AppId),
			sql.Col(tbnAlias2, sql.Id),
		),
		sb.EQ(
			sql.Col(tbnAlias1, db.Group.Type),
			sql.Named(db.Group.Type),
		),
		sb.EQ(
			sql.Col(tbnAlias1, db.Group.AppId),
			sql.Named(db.Group.AppId),
		),
	).OrderBy(
		sql.Col(tbnAlias1, sql.Id),
	).Desc().
		Build()
	return sql.LimitAndOffset(sql.ReplaceAtWithColon(statement))
}

// Update implements group.
func (g *groupImpl) Update() (statement string) {
	sb := sqlbuilder.NewUpdateBuilder()
	statement, _ = sb.Update(db.Group.Tname()).Set(
		sb.Assign(
			db.Group.Name,
			sql.Named(db.Group.Name),
		),
		sb.Assign(
			db.Group.Active,
			sql.Named(db.Group.Active),
		),
	).Where(sb.EQ(sql.Id, sql.Named(sql.Id))).Build()
	return sql.ReplaceAtWithColon(statement)
}

var Group group = &groupImpl{}
