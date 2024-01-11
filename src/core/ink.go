package core

import (
	"github.com/jmoiron/sqlx"
)

type Ink struct {
	Env *Env
	DB  *sqlx.DB
}

func NewInk() *Ink {
	env := loadEnv()
	db := ConnectDB(env)

	return &Ink{
		env,
		db,
	}
}

func (ink *Ink) Close() {
	ink.DB.Close()
}
