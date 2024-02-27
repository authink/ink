package admin

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/i18n"
	"github.com/authink/ink.go/src/migrate"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	inkstone.TestRun(
		"admin",
		&ctx,
		&inkstone.Options{
			Locales: &i18n.Locales,
			Seed: func(appCtx *inkstone.AppContext) {
				migrate.Seed(appCtx)
				if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
					if err = orm.App(appCtx).SaveWithTx(model.NewApp(
						"devtools",
						"123456",
					), tx); err != nil {
						return
					}

					err = orm.Staff(appCtx).SaveWithTx(model.NewStaff(
						"test@huoyijie.cn",
						"123456",
						"11111111111",
						false,
					), tx)
					return
				}); err != nil {
					panic(err)
				}
			},
			SetupAPIGroup: func(apiGroup *gin.RouterGroup) {
				token.SetupTokenGroup(apiGroup)
				SetupAdminGroup(apiGroup, env.AppNameAdmin())
			},
			FinishSetup: authz.SetupEnforcer,
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
