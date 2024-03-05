package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type app struct {
	Name   string
	Secret string
	Active string
}

// Tname implements db.Table.
func (a *app) Tname() string {
	return "s_apps"
}

var _ db.Table = (*app)(nil)

var App app

func init() {
	db.Bind(&App, &models.App{})
}
