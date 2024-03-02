package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/model"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type group interface {
	orm.Inserter[models.Group]
	orm.Updater[models.Group]
	orm.Saver[models.Group]
	orm.Geter[models.Group]
	orm.Counter
	orm.Pager[models.GroupWithApp]
}

type groupImpl app.AppContext

// Count implements group.
func (g *groupImpl) Count(args ...any) (c int, err error) {
	err = count(g.DB, sqls.Group.Count(), &c, args[0])
	return
}

// CountTx implements group.
func (g *groupImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = count(tx, sqls.Group.Count(), &c, args[0])
	return
}

// Get implements group.
// Subtle: this method shadows the method (*DB).Get of groupImpl.DB.
func (g *groupImpl) Get(id int) (group *models.Group, err error) {
	group = new(models.Group)
	err = get(g.DB, group, sqls.Group.Get(), id)
	return
}

// GetTx implements group.
func (g *groupImpl) GetTx(tx *sqlx.Tx, id int) (group *models.Group, err error) {
	group = new(models.Group)
	err = get(tx, group, sqls.Group.GetForUpdate(), id)
	return
}

// Insert implements group.
func (g *groupImpl) Insert(group *models.Group) error {
	return namedExec(g.DB, sqls.Group.Insert(), group, afterInsert)
}

// InsertTx implements group.
func (g *groupImpl) InsertTx(tx *sqlx.Tx, group *models.Group) error {
	return namedExec(tx, sqls.Group.Insert(), group, afterInsert)
}

// PaginationTx implements group.
func (g *groupImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (groups []models.GroupWithApp, err error) {
	err = pagination(tx, sqls.Group.Pagination(), &groups, pager)
	return
}

// Save implements group.
func (g *groupImpl) Save(group *models.Group) error {
	return namedExec(g.DB, sqls.Group.Save(), group, afterSave)
}

// SaveTx implements group.
func (g *groupImpl) SaveTx(tx *sqlx.Tx, group *models.Group) error {
	return namedExec(tx, sqls.Group.Save(), group, afterSave)
}

// Update implements group.
func (g *groupImpl) Update(group *models.Group) error {
	return namedExec(g.DB, sqls.Group.Update(), group, afterUpdate)
}

// UpdateTx implements group.
func (g *groupImpl) UpdateTx(tx *sqlx.Tx, group *models.Group) error {
	return namedExec(tx, sqls.Group.Update(), group, afterUpdate)
}

var _ group = (*groupImpl)(nil)

func Group(appCtx *app.AppContext) group {
	return (*groupImpl)(appCtx)
}
