package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func connectDB(env *Env) *sqlx.DB {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort, env.DbName)

	db, err := sqlx.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
