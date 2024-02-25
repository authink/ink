package admin

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	inkstone.TestMain(
		"admin",
		&ctx,
		&i18n.Locales,
		func(appContext *inkstone.AppContext) {
			migrate.Seed(appContext)
			err := orm.App(appContext).Save(model.NewApp(
				"devtools",
				"123456",
			))
			if err != nil {
				panic(err)
			}
		},
		func(apiGroup *gin.RouterGroup) {
			token.SetupTokenGroup(apiGroup)
			SetupAdminGroup(apiGroup, env.AppNameAdmin())
		},
	)(m)
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
