package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type app interface {
	inkstone.ORM[model.App]
	GetWithTx(int, *sqlx.Tx) (*model.App, error)
}

type appImpl inkstone.AppContext

// Insert implements app.
func (a *appImpl) Insert(app *model.App) (err error) {
	result, err := a.DB.NamedExec(
		sql.App.Insert(),
		app,
	)
	if err != nil {
		return
	}

	err = handleInsertResult(result, &app.Model)
	return
}

// InsertWithTx implements app.
func (*appImpl) InsertWithTx(app *model.App, tx *sqlx.Tx) (err error) {
	result, err := tx.NamedExec(
		sql.App.Insert(),
		app,
	)
	if err != nil {
		return
	}

	err = handleInsertResult(result, &app.Model)
	return
}

// GetWithTx implements app.
func (*appImpl) GetWithTx(id int, tx *sqlx.Tx) (app *model.App, err error) {
	app = new(model.App)
	err = tx.Get(
		app,
		sql.App.GetForUpdate(),
		id,
	)
	return
}

// Delete implements app.
func (*appImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements app.
func (a *appImpl) Find() (apps []model.App, err error) {
	err = a.DB.Select(
		&apps,
		sql.App.Find(),
	)
	return
}

// Get implements app.
func (a *appImpl) Get(id int) (app *model.App, err error) {
	app = new(model.App)
	err = a.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// Save implements app.
func (a *appImpl) Save(app *model.App) (err error) {
	result, err := a.DB.NamedExec(
		sql.App.Save(),
		app,
	)
	if err != nil {
		return
	}

	err = handleSaveResult(result, &app.Model)
	return
}

// SaveWithTx implements app.
func (*appImpl) SaveWithTx(app *model.App, tx *sqlx.Tx) (err error) {
	result, err := tx.NamedExec(
		sql.App.Save(),
		app,
	)
	if err != nil {
		return
	}

	err = handleSaveResult(result, &app.Model)
	return
}

var _ app = (*appImpl)(nil)

func App(appCtx *inkstone.AppContext) app {
	return (*appImpl)(appCtx)
}
