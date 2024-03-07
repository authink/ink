package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/orm/model"
	"github.com/jmoiron/sqlx"
)

type iapp interface {
	orm.Inserter[models.App]
	orm.Updater[models.App]
	orm.Geter[models.App]
	orm.Finder[models.App]
}

type appImpl app.AppContext

// Find implements iapp.
func (a *appImpl) Find(...model.Arg) (apps []models.App, err error) {
	err = orm.Select(a.DB, sqls.App.Find(), &apps, &model.Argument{})
	return
}

// Get implements iapp.
// Subtle: this method shadows the method (*DB).Get of appImpl.DB.
func (a *appImpl) Get(app *models.App) (err error) {
	err = orm.Get(a.DB, sqls.App.Get(), app)
	return
}

// GetTx implements iapp.
func (a *appImpl) GetTx(tx *sqlx.Tx, app *models.App) (err error) {
	err = orm.Get(tx, sqls.App.GetForUpdate(), app)
	return
}

// Insert implements iapp.
func (a *appImpl) Insert(app *models.App) error {
	return orm.NamedInsert(a.DB, sqls.App.Insert(), app)
}

// InsertTx implements iapp.
func (a *appImpl) InsertTx(tx *sqlx.Tx, app *models.App) error {
	return orm.NamedInsert(tx, sqls.App.Insert(), app)
}

// Update implements iapp.
func (a *appImpl) Update(app *models.App) error {
	return orm.NamedUpdate(a.DB, sqls.App.Update(), app)
}

// UpdateTx implements iapp.
func (a *appImpl) UpdateTx(tx *sqlx.Tx, app *models.App) error {
	return orm.NamedUpdate(tx, sqls.App.Update(), app)
}

var _ iapp = (*appImpl)(nil)

func App(appCtx *app.AppContext) iapp {
	return (*appImpl)(appCtx)
}
