package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func createDatabaseUrl(env *Env, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort, env.DbName)

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func connectDB(env *Env) *sqlx.DB {
	databaseUrl := createDatabaseUrl(env, false)

	db, err := sqlx.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatalf("sqlx.Open: %s\n", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(time.Second)

	return db
}
