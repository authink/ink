package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/test"
	"github.com/stretchr/testify/assert"
)

func grantToken(appId int, appSecret, email, password string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &grantReq{
		AppId:     appId,
		AppSecret: appSecret,
		Email:     email,
		Password:  password,
	}

	return test.TestFetch(
		"POST",
		"/api/v1/token/grant",
		reqObj,
		resObj,
	)
}

func TestGrant(t *testing.T) {
	// t.Parallel()
	// go test -v -p 2 # 缺省 -p 参数会根据 cpu 核心数量设置
	var (
		ok                 = []any{http.StatusOK, &grantRes{TokenType: "Bearer", ExpiresIn: 7200}, 100000, "123456", "admin@huoyijie.cn", "123456", &grantRes{}}
		invalidAppId       = []any{http.StatusBadRequest, ext.ERR_INVALID_APP, 999999, "123456", "admin@huoyijie.cn", "123456", &ext.ClientError{}}
		invalidAppSecret   = []any{http.StatusBadRequest, ext.ERR_INVALID_APP, 100000, "1234567", "admin@huoyijie.cn", "123456", &ext.ClientError{}}
		invalidEmailFormat = []any{http.StatusBadRequest, nil, 100000, "123456", "admin", "123456", &ext.ClientError{}}
		invalidEmail       = []any{http.StatusBadRequest, ext.ERR_INVALID_ACCOUNT, 100000, "123456", "admin1@huoyijie.cn", "123456", &ext.ClientError{}}
		invalidPassword    = []any{http.StatusBadRequest, ext.ERR_INVALID_ACCOUNT, 100000, "123456", "admin@huoyijie.cn", "1234567", &ext.ClientError{}}
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
			resObj := tc[6].(*grantRes)
			assert.Equal(t, tc[1].(*grantRes).TokenType, resObj.TokenType)
			assert.Equal(t, tc[1].(*grantRes).ExpiresIn, resObj.ExpiresIn)
			assert.NotEmpty(t, resObj.AccessToken)
			assert.NotEmpty(t, resObj.RefreshToken)

		case http.StatusBadRequest:
			if tc[1] != nil {
				resObj := tc[6].(*ext.ClientError)
				assert.Equal(t, tc[1].(*ext.ClientError).Code, resObj.Code)
			}
		}
	}
}
