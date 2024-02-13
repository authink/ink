package service

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

const (
	APP_ADMIN_DEV string = "admin.dev"
)

type appService interface {
	SaveApp(*model.App) error
	SaveAppWithTx(*model.App, *sqlx.Tx) error
	GetApp(int) (*model.App, error)
}

type AppService core.Ink

// GetApp implements appService.
func (as *AppService) GetApp(id int) (app *model.App, err error) {
	app = &model.App{}
	err = as.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// SaveApp implements appService.
func (as *AppService) SaveApp(app *model.App) (err error) {
	_, err = as.DB.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

// SaveAppWithTx implements appService.
func (*AppService) SaveAppWithTx(app *model.App, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.App.Insert(),
		app,
	)
	return
}

var _ appService = (*AppService)(nil)
