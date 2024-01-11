package core

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type appService interface {
	SaveApp(*model.App) error
	GetApp(id int) (*model.App, error)
}

// GetApp implements appService.
func (ink *Ink) GetApp(id int) (app *model.App, err error) {
	app = &model.App{}
	err = ink.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// SaveApp implements appService.
func (ink *Ink) SaveApp(*model.App) error {
	panic("unimplemented")
}

var _ appService = (*Ink)(nil)
