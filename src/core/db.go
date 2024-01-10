package core

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CreateDatabaseUrl(env *Env, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort, env.DbName)

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func ConnectDB(env *Env) *sqlx.DB {
	databaseUrl := CreateDatabaseUrl(env, false)

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

func TxEnd(tx *sqlx.Tx, err error) {
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
