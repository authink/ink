package token

import (
	"context"
	"testing"

	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	inkstone.TestMain(
		"token",
		&ctx,
		&i18n.Locales,
		migrate.Seed,
		func(apiGroup *gin.RouterGroup) {
			SetupTokenGroup(apiGroup)
		},
	)(m)
}
