package token

import (
	"context"
	"fmt"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router/common"
	"github.com/authink/ink.go/src/test"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	env := core.LoadEnv()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, "token")
	defer core.CreateDB(env)()

	ink := core.NewInkWith(env)
	defer ink.Close()

	router, gApi := common.SetupRouter(ink)
	SetupTokenGroup(gApi)

	test.Main(&ctx, ink, router)(m)
}
