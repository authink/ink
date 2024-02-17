package core

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDatabaseUrl(env *Env, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort, env.DbName)

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func ConnectDB(env *Env) *sqlx.DB {
	databaseUrl := GetDatabaseUrl(env, false)

	db, err := sqlx.Open("mysql", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(int(env.DbMaxOpenConns))
	db.SetMaxIdleConns(int(env.DbMaxIdleConns))
	db.SetConnMaxLifetime(time.Duration(env.DbConnMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(env.DbConnMaxIdleTime) * time.Second)

	return db
}

func CreateDB(env *Env) func() {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/", env.DbUser, env.DbPasswd, env.DbHost, env.DbPort)

	db, err := sqlx.Open("mysql", databaseUrl)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", env.DbName))
	if err != nil {
		db.Close()
		panic(err)
	}

	return func() {
		defer db.Close()
		_, err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", env.DbName))
		if err != nil {
			panic(err)
		}
	}
}
