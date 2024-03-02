package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type app interface {
	orm.Inserter[model.App]
	orm.Saver[model.App]
	orm.Updater[model.App]
	orm.Geter[model.App]
	orm.Finder[model.App]
}

type appImpl a.AppContext

// Find implements app.
func (a *appImpl) Find(...any) (apps []model.App, err error) {
	err = a.DB.Select(
		&apps,
		sql.App.Find(),
	)
	return
}

// Get implements app.
// Subtle: this method shadows the method (*DB).Get of appImpl.DB.
func (a *appImpl) Get(id int) (app *model.App, err error) {
	app = new(model.App)
	err = a.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// GetTx implements app.
func (a *appImpl) GetTx(tx *sqlx.Tx, id int) (app *model.App, err error) {
	app = new(model.App)
	err = tx.Get(
		app,
		sql.App.GetForUpdate(),
		id,
	)
	return
}

// Insert implements app.
func (a *appImpl) Insert(app *model.App) error {
	return namedExec(a.DB, sql.App.Insert(), app, handleInsertResult)
}

// InsertTx implements app.
func (a *appImpl) InsertTx(tx *sqlx.Tx, app *model.App) error {
	return namedExec(tx, sql.App.Insert(), app, handleInsertResult)
}

// Save implements app.
func (a *appImpl) Save(app *model.App) error {
	return namedExec(a.DB, sql.App.Save(), app, handleSaveResult)
}

// SaveTx implements app.
func (a *appImpl) SaveTx(tx *sqlx.Tx, app *model.App) error {
	return namedExec(tx, sql.App.Save(), app, handleSaveResult)
}

// Update implements app.
func (a *appImpl) Update(app *model.App) error {
	return namedExec(a.DB, sql.App.Update(), app, nil)
}

// UpdateTx implements app.
func (a *appImpl) UpdateTx(tx *sqlx.Tx, app *model.App) error {
	return namedExec(tx, sql.App.Update(), app, nil)
}

var _ app = (*appImpl)(nil)

func App(appCtx *a.AppContext) app {
	return (*appImpl)(appCtx)
}
