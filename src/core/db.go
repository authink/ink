package core

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CreateDatabaseUrl(env *Env, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort, env.DbName)

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

	db.SetMaxOpenConns(int(env.DbMaxOpenConns))
	db.SetMaxIdleConns(int(env.DbMaxIdleConns))
	db.SetConnMaxLifetime(time.Duration(env.DbConnMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(env.DbConnMaxIdleTime) * time.Second)

	return db
}

func TxEnd(tx *sqlx.Tx, err error) {
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
