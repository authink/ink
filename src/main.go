package main

import (
	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/router"
	"github.com/authink/inkstone"
)

func main() {
	inkstone.Main(
		&i18n.Locales,
		migrate.Seed,
		router.SetupRouter,
	)
}
