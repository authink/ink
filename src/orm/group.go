package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/orm/model"
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
	err = orm.Count(g.DB, sqls.Group.Count(), &c, args[0])
	return
}

// CountTx implements group.
func (g *groupImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = orm.Count(tx, sqls.Group.Count(), &c, args[0])
	return
}

// Get implements group.
// Subtle: this method shadows the method (*DB).Get of groupImpl.DB.
func (g *groupImpl) Get(id int) (group *models.Group, err error) {
	group = &models.Group{}
	err = orm.Get(g.DB, group, sqls.Group.Get(), id)
	return
}

// GetTx implements group.
func (g *groupImpl) GetTx(tx *sqlx.Tx, id int) (group *models.Group, err error) {
	group = &models.Group{}
	err = orm.Get(tx, group, sqls.Group.GetForUpdate(), id)
	return
}

// Insert implements group.
func (g *groupImpl) Insert(group *models.Group) error {
	return orm.NamedInsert(g.DB, sqls.Group.Insert(), group)
}

// InsertTx implements group.
func (g *groupImpl) InsertTx(tx *sqlx.Tx, group *models.Group) error {
	return orm.NamedInsert(tx, sqls.Group.Insert(), group)
}

// PaginationTx implements group.
func (g *groupImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (groups []models.GroupWithApp, err error) {
	err = orm.Pagination(tx, sqls.Group.Pagination(), &groups, pager)
	return
}

// Save implements group.
func (g *groupImpl) Save(group *models.Group) error {
	return orm.NamedSave(g.DB, sqls.Group.Save(), group)
}

// SaveTx implements group.
func (g *groupImpl) SaveTx(tx *sqlx.Tx, group *models.Group) error {
	return orm.NamedSave(tx, sqls.Group.Save(), group)
}

// Update implements group.
func (g *groupImpl) Update(group *models.Group) error {
	return orm.NamedUpdate(g.DB, sqls.Group.Update(), group)
}

// UpdateTx implements group.
func (g *groupImpl) UpdateTx(tx *sqlx.Tx, group *models.Group) error {
	return orm.NamedUpdate(tx, sqls.Group.Update(), group)
}

var _ group = (*groupImpl)(nil)

func Group(appCtx *app.AppContext) group {
	return (*groupImpl)(appCtx)
}
