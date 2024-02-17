package admin

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router/common"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/ink.go/src/test"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	ink := core.NewInk()
	defer ink.Close()

	router, gApi := common.SetupRouter(ink)
	token.SetupTokenGroup(gApi)
	SetupAdminGroup(gApi, ink.Env.AppNameAdmin)

	test.Main(&ctx, ink, router)(m)
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
		"POST",
		"token/grant",
		reqObj,
		resObj,
		"",
	)
}
