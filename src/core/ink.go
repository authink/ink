package core

import (
	myI18n "github.com/authink/ink.go/src/i18n"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Ink struct {
	Env    *Env
	DB     *sqlx.DB
	Bundle *i18n.Bundle
}

func NewInk() *Ink {
	return NewInkWith(LoadEnv())
}

func NewInkWith(env *Env) *Ink {
	db := ConnectDB(env)
	return &Ink{
		env,
		db,
		myI18n.NewBundle(),
	}
}

func (ink *Ink) Close() {
	ink.DB.Close()
}
