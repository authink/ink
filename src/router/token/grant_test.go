package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/inkstone"
	"github.com/stretchr/testify/assert"
)

func grantToken(appId int, appSecret, email, password string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &GrantReq{
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

func TestGrant(t *testing.T) {
	// t.Parallel()
	// go test -v -p 2 # 缺省 -p 参数会根据 cpu 核心数量设置
	var (
		ok                 = []any{http.StatusOK, &GrantRes{TokenType: "Bearer", ExpiresIn: 7200}, 100000, "123456", "admin@huoyijie.cn", "123456", new(GrantRes)}
		invalidAppId       = []any{http.StatusBadRequest, errors.ERR_INVALID_APP, 999999, "123456", "admin@huoyijie.cn", "123456", new(inkstone.ClientError)}
		invalidAppSecret   = []any{http.StatusBadRequest, errors.ERR_INVALID_APP, 100000, "1234567", "admin@huoyijie.cn", "123456", new(inkstone.ClientError)}
		invalidEmailFormat = []any{http.StatusBadRequest, nil, 100000, "123456", "admin", "123456", new(inkstone.ClientError)}
		invalidEmail       = []any{http.StatusBadRequest, errors.ERR_INVALID_ACCOUNT, 100000, "123456", "admin1@huoyijie.cn", "123456", new(inkstone.ClientError)}
		invalidPassword    = []any{http.StatusBadRequest, errors.ERR_INVALID_ACCOUNT, 100000, "123456", "admin@huoyijie.cn", "1234567", new(inkstone.ClientError)}
	)

	var tests = []any{
		ok,
		invalidAppId,
		invalidAppSecret,
		invalidEmailFormat,
		invalidEmail,
		invalidPassword,
	}

	for i, test := range tests {
		t.Logf("TestGrant: [%d]\n", i)
		tc := test.([]any)

		w, _ := grantToken(tc[2].(int), tc[3].(string), tc[4].(string), tc[5].(string), tc[6])
		assert.Equal(t, tc[0].(int), w.Code)

		switch tc[0].(int) {
		case http.StatusOK:
			resObj := tc[6].(*GrantRes)
			assert.Equal(t, tc[1].(*GrantRes).TokenType, resObj.TokenType)
			assert.Equal(t, tc[1].(*GrantRes).ExpiresIn, resObj.ExpiresIn)
			assert.NotEmpty(t, resObj.AccessToken)
			assert.NotEmpty(t, resObj.RefreshToken)

		case http.StatusBadRequest:
			if tc[1] != nil {
				resObj := tc[6].(*inkstone.ClientError)
				assert.Equal(t, tc[1].(*inkstone.ClientError).Code, resObj.Code)
			}
		}
	}
}
