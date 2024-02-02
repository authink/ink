package core

import (
	"github.com/jmoiron/sqlx"
)

type TxFunc func(tx *sqlx.Tx) error

type Ink struct {
	db  *sqlx.DB
	Env *Env
}

func NewInk() *Ink {
	env := loadEnv()
	db := ConnectDB(env)

	return &Ink{
		db,
		env,
	}
}

func (ink *Ink) Close() {
	ink.db.Close()
}

func (ink *Ink) Transaction(txFunc TxFunc) (err error) {
	tx := ink.db.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
