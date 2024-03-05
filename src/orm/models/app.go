package models

import (
	"github.com/authink/inkstone/orm/model"
	"github.com/authink/inkstone/util"
)

// @model
// @db s_apps
type App struct {
	model.Base
	Name   string
	Secret string
	Active bool
}

func NewApp(name, secret string) *App {
	return &App{
		Name:   name,
		Secret: util.Sha256(secret),
		Active: true,
	}
}

func (app *App) Reset(secret string) {
	app.Secret = util.Sha256(secret)
}
