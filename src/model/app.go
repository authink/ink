package model

import (
	"github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/util"
)

type App struct {
	orm.Model
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
