package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if inkCWD := getInkCWD(); inkCWD != "" {
		if err := os.Chdir(inkCWD); err != nil {
			panic(err)
		}
	}

	inkENV := getInkENV()

	files := []string{
		fmt.Sprintf(".env.%s.local", inkENV),
		".env.local",
		fmt.Sprintf(".env.%s", inkENV),
		".env",
	}

	var existFiles []string
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			existFiles = append(existFiles, file)
		}
	}

	if len(existFiles) > 0 {
		if err := godotenv.Load(existFiles...); err != nil {
			panic(err)
		}
	}
}

const (
	DEVELOPMENT string = "dev"
	TEST        string = "test"
	PRODUCTION  string = "prod"
)

type Env struct {
	InkENV               string
	InkCWD               string
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

func getInkENV() string {
	inkENV := DEVELOPMENT
	getString("INK_ENV", &inkENV)

	if !(inkENV == DEVELOPMENT || inkENV == TEST || inkENV == PRODUCTION) {
		panic(fmt.Sprintf("Invalid INK_ENV %s", inkENV))
	}
	return inkENV
}

func getInkCWD() (inkCWD string) {
	getString("INK_CWD", &inkCWD)
	return
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
	inkENV := getInkENV()
	inkCWD := getInkCWD()
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
		inkENV,
		inkCWD,
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
