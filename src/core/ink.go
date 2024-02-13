package core

import (
	myI18n "github.com/authink/ink.go/src/i18n"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TxFunc func(tx *sqlx.Tx) error

type Ink struct {
	Env    *Env
	DB     *sqlx.DB
	Bundle *i18n.Bundle
}

func NewInk() *Ink {
	env := loadEnv()
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

func (ink *Ink) Transaction(txFunc TxFunc) (err error) {
	tx := ink.DB.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
