package token

import (
	"context"
	"fmt"
	"testing"

	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/inkstone"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	env := inkstone.LoadEnv()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, "token")

	defer inkstone.CreateDB(
		env.DbUser,
		env.DbPasswd,
		env.DbName,
		env.DbHost,
		env.DbPort,
	)()

	app := inkstone.NewAppContextWithEnv(&i18n.Locales, env)
	defer app.Close()

	router, apiGroup := inkstone.SetupRouter(app)
	SetupTokenGroup(apiGroup)

	inkstone.TestMain(&ctx, app, router, migrate.Seed)(m)
}
