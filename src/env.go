package main

import (
	"fmt"
	"log"
	"os"
)

type Env struct {
	Host              string
	Port              uint16
	ShutdownTimeout   uint16
	DbHost            string
	DbPort            uint16
	DbUser            string
	DbPasswd          string
	DbName            string
	DbMaxOpenConns    uint16
	DbMaxIdleConns    uint16
	DbConnMaxLifeTime uint16
	DbConnMaxIdleTime uint16
}

func getUint16(key string, value *uint16) {
	if v := os.Getenv(key); len(v) > 0 {
		if _, err := fmt.Sscanf(v, "%d", value); err != nil {
			log.Fatal(err)
		}
	}
}

func getString(key string, value *string) {
	if v := os.Getenv(key); len(v) > 0 {
		*value = v
	}
}

func loadEnv() *Env {
	host := "localhost"
	port := uint16(8080)
	shutdownTimeout := uint16(5)

	getString("HOST", &host)
	getUint16("PORT", &port)
	getUint16("SHUTDOWN_TIMEOUT", &shutdownTimeout)

	dbHost := "localhost"
	dbPort := uint16(3306)
	dbUser := "root"
	dbPasswd := "root"
	dbName := "ink"

	getString("DB_HOST", &dbHost)
	getUint16("DB_PORT", &dbPort)
	getString("DB_USER", &dbUser)
	getString("DB_PASSWORD", &dbPasswd)
	getString("DB_NAME", &dbName)

	dbMaxOpenConns := uint16(20)
	dbMaxIdleConns := uint16(10)
	dbConnMaxLifeTime := uint16(3600)
	dbConnMaxIdleTime := uint16(300)
	getUint16("DB_MAX_OPEN_CONNS", &dbMaxOpenConns)
	getUint16("DB_MAX_IDLE_CONNS", &dbMaxIdleConns)
	getUint16("DB_CONN_MAX_LIFE_TIME", &dbConnMaxLifeTime)
	getUint16("DB_CONN_MAX_IDLE_TIME", &dbConnMaxIdleTime)

	return &Env{
		host,
		port,
		shutdownTimeout,
		dbHost,
		dbPort,
		dbUser,
		dbPasswd,
		dbName,
		dbMaxOpenConns,
		dbMaxIdleConns,
		dbConnMaxLifeTime,
		dbConnMaxIdleTime,
	}
}
