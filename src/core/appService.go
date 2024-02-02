package core

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

type appService interface {
	SaveApp(*model.App) error
	SaveAppWithTx(*model.App, *sqlx.Tx) error
	GetApp(int) (*model.App, error)
}

// GetApp implements appService.
func (ink *Ink) GetApp(id int) (app *model.App, err error) {
	app = &model.App{}
	err = ink.db.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// SaveApp implements appService.
func (ink *Ink) SaveApp(app *model.App) (err error) {
	_, err = ink.db.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

// SaveAppWithTx implements appService.
func (ink *Ink) SaveAppWithTx(app *model.App, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

var _ appService = (*Ink)(nil)
