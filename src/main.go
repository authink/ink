package main

import (
	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/migrate"
	"github.com/authink/ink/src/web/i18n"
	"github.com/authink/ink/src/web/router"
	"github.com/authink/stone"
	"github.com/authink/stone/app"
)

func main() {
	stone.Run(&app.Options{
		Locales:       &i18n.Locales,
		Seed:          migrate.Seed,
		SetupApiGroup: router.SetupRouter,
		FinishSetup:   authz.SetupEnforcer,
	})
}
