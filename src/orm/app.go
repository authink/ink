package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type iapp interface {
	orm.Inserter[models.App]
	orm.Saver[models.App]
	orm.Updater[models.App]
	orm.Geter[models.App]
	orm.Finder[models.App]
}

type appImpl app.AppContext

// Find implements iapp.
func (a *appImpl) Find(args ...any) (apps []models.App, err error) {
	err = orm.Select(a.DB, &apps, sqls.App.Find(), args...)
	return
}

// Get implements iapp.
// Subtle: this method shadows the method (*DB).Get of appImpl.DB.
func (a *appImpl) Get(id int) (app *models.App, err error) {
	app = new(models.App)
	err = orm.Get(a.DB, app, sqls.App.Get(), id)
	return
}

// GetTx implements iapp.
func (a *appImpl) GetTx(tx *sqlx.Tx, id int) (app *models.App, err error) {
	app = new(models.App)
	err = orm.Get(tx, app, sqls.App.GetForUpdate(), id)
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

// Save implements iapp.
func (a *appImpl) Save(app *models.App) error {
	return orm.NamedSave(a.DB, sqls.App.Save(), app)
}

// SaveTx implements iapp.
func (a *appImpl) SaveTx(tx *sqlx.Tx, app *models.App) error {
	return orm.NamedSave(tx, sqls.App.Save(), app)
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
