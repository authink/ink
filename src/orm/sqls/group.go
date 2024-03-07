package sqls

import (
	"fmt"

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
func (g *groupImpl) Pagination() string {
	
	return fmt.Sprintf("SELECT g.id, g.created_at, g.updated_at, g.name, g.type, g.app_id, a.name app_name, g.active FROM %s g, %s a WHERE g.app_id = a.id AND g.type = :type AND g.app_id = :app_id ORDER BY g.id DESC LIMIT :limit OFFSET :offset", db.Group.Tname(), db.App.Tname())
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
