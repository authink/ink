package main

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/web/i18n"
	"github.com/authink/ink.go/src/web/router"
	"github.com/authink/inkstone"
	"github.com/authink/inkstone/app"
)

func main() {
	inkstone.Run(&app.Options{
		Locales:       &i18n.Locales,
		Seed:          migrate.Seed,
		SetupApiGroup: router.SetupRouter,
		FinishSetup:   authz.SetupEnforcer,
	})
}
