package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func createSourceUrl(ink *Ink) string {
	return fmt.Sprintf("file://%s", ink.env.DbMigrateFileSource)
}

func migrateSchema(ink *Ink) {
	sourceUrl := createSourceUrl(ink)
	databaseUrl := createDatabaseUrl(ink.env, true)

	m, err := migrate.New(
		sourceUrl,
		databaseUrl,
	)

	if err != nil {
		log.Fatalf("Migrate Init: %s,%s,%s\n", sourceUrl, databaseUrl, err)
	}

	if err = m.Up(); err != nil {
		log.Fatalf("Migrate Down: %s\n", err)
	}
}
