package admin

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/migrate"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/i18n"
	"github.com/authink/ink/src/web/router/token"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/test"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	test.Run(
		"admin",
		&ctx,
		&app.Options{
			Locales: &i18n.Locales,
			Seed: func(appCtx *app.AppContext) {
				migrate.Seed(appCtx)
				if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
					if err = orm.App(appCtx).InsertTx(tx, models.NewApp(
						"devtools",
						"123456",
					)); err != nil {
						return
					}

					if err = orm.Staff(appCtx).InsertTx(tx, models.NewStaff(
						"test@huoyijie.cn",
						"123456",
						"11111111111",
						false,
					)); err != nil {
						return
					}

					err = orm.Group(appCtx).InsertTx(tx, models.NewGroup(
						"developer",
						1,
						100000,
					))
					return
				}); err != nil {
					panic(err)
				}
			},
			SetupApiGroup: func(apiGroup *gin.RouterGroup) {
				token.SetupTokenGroup(apiGroup)
				SetupAdminGroup(apiGroup, envs.AppNameAdmin())
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

	return test.Fetch(
		ctx,
		http.MethodPost,
		"token/grant",
		reqObj,
		resObj,
		"",
	)
}
