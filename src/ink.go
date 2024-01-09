package main

import "github.com/jmoiron/sqlx"

type Ink struct {
	env *Env
	db  *sqlx.DB
}

func newInk() *Ink {
	env := loadEnv()
	db := connectDB(env)

	return &Ink{
		env,
		db,
	}
}

func (ink *Ink) Close() {
	ink.db.Close()
}
