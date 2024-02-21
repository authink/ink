package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type app interface {
	inkstone.ORM[model.App]
}

type appImpl inkstone.AppContext

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
	app = &model.App{}
	err = a.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// Save implements app.
func (a *appImpl) Save(app *model.App) (err error) {
	_, err = a.DB.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

// SaveWithTx implements app.
func (*appImpl) SaveWithTx(app *model.App, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

var _ app = (*appImpl)(nil)

func App(appContext *inkstone.AppContext) app {
	return (*appImpl)(appContext)
}
