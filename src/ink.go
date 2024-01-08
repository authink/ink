package main

import "github.com/jmoiron/sqlx"

type Ink struct {
	env *Env
	db  *sqlx.DB
}
