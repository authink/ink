package admin

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	myEnv "github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	env := inkstone.LoadEnv()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, "admin")
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
	token.SetupTokenGroup(apiGroup)
	SetupAdminGroup(apiGroup, myEnv.AppNameAdmin())

	inkstone.TestMain(&ctx, app, router, migrate.Seed)(m)
}

func grantToken(appId int, appSecret, email, password string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &token.GrantReq{
		AppId:     appId,
		AppSecret: appSecret,
		Email:     email,
		Password:  password,
	}

	return inkstone.TestFetch(
		ctx,
		"POST",
		"token/grant",
		reqObj,
		resObj,
		"",
	)
}
