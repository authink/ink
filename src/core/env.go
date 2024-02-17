package core

import (
	"fmt"
	"os"
)

type Env struct {
	SecretKey            string
	AccessTokenDuration  uint16
	RefreshTokenDuration uint16
	Host                 string
	Port                 uint16
	ShutdownTimeout      uint16
	DbHost               string
	DbPort               uint16
	DbUser               string
	DbPasswd             string
	DbName               string
	DbMaxOpenConns       uint16
	DbMaxIdleConns       uint16
	DbConnMaxLifeTime    uint16
	DbConnMaxIdleTime    uint16
	DbMigrateFileSource  string
	AppNameAdmin         string
	BasePath             string
}

func getUint16(key string, value *uint16) {
	if v := os.Getenv(key); len(v) > 0 {
		if _, err := fmt.Sscanf(v, "%d", value); err != nil {
			panic(err)
		}
	}
}

func getString(key string, value *string) {
	if v := os.Getenv(key); len(v) > 0 {
		*value = v
	}
}

func LoadEnv() *Env {
	secretKey := "your-secret-key"
	accessTokenDuration := uint16(7200)
	refreshTokenDuration := uint16(7 * 24)
	host := "localhost"
	port := uint16(8080)
	shutdownTimeout := uint16(5)

	getString("SECRET_KEY", &secretKey)
	getUint16("ACCESS_TOKEN_DURATION", &accessTokenDuration)
	getUint16("REFRESH_TOKEN_DURATION", &refreshTokenDuration)

	getString("HOST", &host)
	getUint16("PORT", &port)
	getUint16("SHUTDOWN_TIMEOUT", &shutdownTimeout)

	dbHost := "localhost"
	dbPort := uint16(3306)
	dbUser := "root"
	dbPasswd := ""
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

	dbMigrateFileSource := "ink.schema/migrations"
	getString("DB_MIGRATE_FILE_SOURCE", &dbMigrateFileSource)

	appNameAdmin := "admin.dev"
	getString("APP_NAME_ADMIN", &appNameAdmin)

	basePath := "api/v1"
	getString("BASE_PATH", &basePath)

	return &Env{
		secretKey,
		accessTokenDuration,
		refreshTokenDuration,
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
		dbMigrateFileSource,
		appNameAdmin,
		basePath,
	}
}
