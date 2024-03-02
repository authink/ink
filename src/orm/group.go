package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type GroupPageArg struct {
	orm.PageArgs
	Type  int
	AppId int `db:"app_id"`
}

type group interface {
	orm.Inserter[model.Group]
	orm.Updater[model.Group]
	orm.Saver[model.Group]
	orm.Geter[model.Group]
	orm.Counter
	orm.Pager[model.GroupWithApp]
}

type groupImpl a.AppContext

// Count implements group.
func (g *groupImpl) Count(args ...any) (c int, err error) {
	stmt, err := g.DB.PrepareNamed(sql.Group.Count())
	if err != nil {
		return
	}
	err = stmt.Get(&c, args[0])
	return
}

// CountTx implements group.
func (g *groupImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	stmt, err := tx.PrepareNamed(sql.Group.Count())
	if err != nil {
		return
	}
	err = stmt.Get(&c, args[0])
	return
}

// Get implements group.
// Subtle: this method shadows the method (*DB).Get of groupImpl.DB.
func (g *groupImpl) Get(id int) (group *model.Group, err error) {
	group = new(model.Group)
	err = g.DB.Get(
		group,
		sql.Group.Get(),
		id,
	)
	return
}

// GetTx implements group.
func (g *groupImpl) GetTx(tx *sqlx.Tx, id int) (group *model.Group, err error) {
	group = new(model.Group)
	err = tx.Get(
		group,
		sql.Group.GetForUpdate(),
		id,
	)
	return
}

// Insert implements group.
func (g *groupImpl) Insert(group *model.Group) error {
	return namedExec(g.DB, sql.Group.Insert(), group, handleInsertResult)
}

// InsertTx implements group.
func (g *groupImpl) InsertTx(tx *sqlx.Tx, group *model.Group) error {
	return namedExec(tx, sql.Group.Insert(), group, handleInsertResult)
}

// PaginationTx implements group.
func (g *groupImpl) PaginationTx(tx *sqlx.Tx, page orm.Page) (groups []model.GroupWithApp, err error) {
	stmt, err := tx.PrepareNamed(sql.Group.Pagination())
	if err != nil {
		return
	}
	err = stmt.Select(&groups, page)
	return
}

// Save implements group.
func (g *groupImpl) Save(group *model.Group) error {
	return namedExec(g.DB, sql.Group.Save(), group, handleSaveResult)
}

// SaveTx implements group.
func (g *groupImpl) SaveTx(tx *sqlx.Tx, group *model.Group) error {
	return namedExec(tx, sql.Group.Save(), group, handleSaveResult)
}

// Update implements group.
func (g *groupImpl) Update(group *model.Group) error {
	return namedExec(g.DB, sql.Group.Update(), group, nil)
}

// UpdateTx implements group.
func (g *groupImpl) UpdateTx(tx *sqlx.Tx, group *model.Group) error {
	return namedExec(tx, sql.Group.Update(), group, nil)
}

var _ group = (*groupImpl)(nil)

func Group(appCtx *a.AppContext) group {
	return (*groupImpl)(appCtx)
}
