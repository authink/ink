package migrate

import (
	"fmt"

	"github.com/authink/ink.go/src/core"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func createSourceUrl(ink *core.Ink) string {
	return fmt.Sprintf("file://%s", ink.Env.DbMigrateFileSource)
}

func Schema(ink *core.Ink, direction string) {
	if direction != "up" && direction != "down" {
		panic(fmt.Errorf("migrate: unkwon direction %s", direction))
	}

	sourceUrl := createSourceUrl(ink)
	databaseUrl := core.GetDatabaseUrl(ink.Env, true)

	m, err := migrate.New(
		sourceUrl,
		databaseUrl,
	)

	if err != nil {
		panic(err)
	}

	switch direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	}

	if err != nil {
		panic(err)
	}
}
