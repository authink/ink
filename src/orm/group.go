package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type group interface {
	inkstone.ORM[model.Group]
	GetWithTx(int, *sqlx.Tx) (*model.Group, error)
	CountWithTx(gtype, appId int, tx *sqlx.Tx) (int, error)
	PaginationWithTx(gtype, appId, offset, limit int, tx *sqlx.Tx) ([]model.GroupWithApp, error)
}

type groupImpl inkstone.AppContext

// GetWithTx implements group.
func (g *groupImpl) GetWithTx(id int, tx *sqlx.Tx) (group *model.Group, err error) {
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

// InsertWithTx implements group.
func (g *groupImpl) InsertWithTx(group *model.Group, tx *sqlx.Tx) error {
	return namedExec(tx, sql.Group.Insert(), group, handleInsertResult)
}

// CountWithTx implements group.
func (*groupImpl) CountWithTx(gtype, appId int, tx *sqlx.Tx) (c int, err error) {
	err = tx.Get(&c, sql.Group.Count(), gtype, appId)
	return
}

// PaginationWithTx implements group.
func (*groupImpl) PaginationWithTx(gtype, appId, offset, limit int, tx *sqlx.Tx) (groups []model.GroupWithApp, err error) {
	err = tx.Select(
		&groups,
		sql.Group.Pagination(),
		gtype,
		appId,
		limit,
		offset,
	)
	return
}

// Delete implements group.
func (*groupImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements group.
func (*groupImpl) Find() ([]model.Group, error) {
	panic("unimplemented")
}

// Get implements group.
func (*groupImpl) Get(int) (*model.Group, error) {
	panic("unimplemented")
}

// Save implements group.
func (g *groupImpl) Save(group *model.Group) error {
	return namedExec(g.DB, sql.Group.Save(), group, handleSaveResult)
}

// SaveWithTx implements group.
func (*groupImpl) SaveWithTx(group *model.Group, tx *sqlx.Tx) error {
	return namedExec(tx, sql.Group.Save(), group, handleSaveResult)
}

var _ group = (*groupImpl)(nil)

func Group(appCtx *inkstone.AppContext) group {
	return (*groupImpl)(appCtx)
}
