package main

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/router"
	"github.com/authink/inkstone"
)

func main() {
	inkstone.Run(&inkstone.Options{
		Locales:       &i18n.Locales,
		Seed:          migrate.Seed,
		SetupAPIGroup: router.SetupRouter,
		FinishSetup:   authz.SetupEnforcer,
	})
}
