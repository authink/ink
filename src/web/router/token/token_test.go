package token

import (
	"context"
	"testing"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/migrate"
	"github.com/authink/ink/src/web/i18n"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/test"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	test.Run(
		"token",
		&ctx,
		&app.Options{
			Locales: &i18n.Locales,
			Seed:    migrate.Seed,
			SetupApiGroup: func(apiGroup *gin.RouterGroup) {
				SetupTokenGroup(apiGroup)
			},
			FinishSetup: authz.SetupEnforcer,
		},
	)(m)
}
