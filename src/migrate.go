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

func migrateSchema(ink *Ink, direction string) {
	if direction != "up" && direction != "down" {
		log.Fatalf("Migrate: unkwon direction %s\n", direction)
	}

	sourceUrl := createSourceUrl(ink)
	databaseUrl := createDatabaseUrl(ink.env, true)

	m, err := migrate.New(
		sourceUrl,
		databaseUrl,
	)

	if err != nil {
		log.Fatalf("Migrate Init: %s,%s,%s\n", sourceUrl, databaseUrl, err)
	}

	switch direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	}

	if err != nil {
		log.Fatalf("Migrate: %s\n", err)
	}
}
