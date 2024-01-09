package models

import "github.com/authink/ink.go/src/util"

type App struct {
	Model
	Name   string
	Secret string
}

func NewApp(name, secret string) *App {
	return &App{
		Name:   name,
		Secret: util.Sha256(secret),
	}
}
