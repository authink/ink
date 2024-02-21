package model

import (
	"github.com/authink/inkstone"
)

type App struct {
	inkstone.Model
	Name   string
	Secret string
	Active bool
}

func NewApp(name, secret string) *App {
	return &App{
		Name:   name,
		Secret: inkstone.Sha256(secret),
		Active: true,
	}
}
