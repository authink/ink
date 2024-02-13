package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

type app interface {
	Save(*model.App) error
	SaveWithTx(*model.App, *sqlx.Tx) error
	Get(int) (*model.App, error)
}

type appImpl core.Ink

// Get implements app.
func (as *appImpl) Get(id int) (app *model.App, err error) {
	app = &model.App{}
	err = as.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// Save implements app.
func (as *appImpl) Save(app *model.App) (err error) {
	_, err = as.DB.NamedExec(
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

func App(ink *core.Ink) app {
	return (*appImpl)(ink)
}
