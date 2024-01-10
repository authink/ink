package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/ext"
	"github.com/stretchr/testify/assert"
)

func grantToken(appId int, appSecret, email, password string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &reqGrant{
		AppId:     appId,
		AppSecret: appSecret,
		Email:     email,
		Password:  password,
	}

	return ext.TestFetch(
		"POST",
		"/token/grant",
		reqObj,
		resObj,
		SetupTokenGroup,
	)
}

func TestGrantToken(t *testing.T) {
	resObj := &resGrant{}
	w, err := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Bearer", resObj.TokenType)
	assert.Equal(t, 7200, resObj.ExpiresIn)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)
}

func TestGrantTokenInvalidAppId(t *testing.T) {
	resObj := &ext.ClientError{}
	w, err := grantToken(999999, "123456", "admin@huoyijie.cn", "123456", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, ext.ERR_CLI_INVALID_APP.(*ext.ClientError).Code, resObj.Code)
}

func TestGrantTokenInvalidAppSecret(t *testing.T) {
	resObj := &ext.ClientError{}
	w, err := grantToken(100000, "1234567", "admin@huoyijie.cn", "123456", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, ext.ERR_CLI_INVALID_APP.(*ext.ClientError).Code, resObj.Code)
}

func TestGrantTokenInvalidEmailFormat(t *testing.T) {
	resObj := &ext.ClientError{}
	w, err := grantToken(100000, "123456", "admin", "123456", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGrantTokenInvalidEmail(t *testing.T) {
	resObj := &ext.ClientError{}
	w, err := grantToken(100000, "123456", "admin1@huoyijie.cn", "123456", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, ext.ERR_CLI_INVALID_ACCOUNT.(*ext.ClientError).Code, resObj.Code)
}

func TestGrantTokenInvalidPassword(t *testing.T) {
	resObj := &ext.ClientError{}
	w, err := grantToken(100000, "123456", "admin@huoyijie.cn", "1234567", resObj)

	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, ext.ERR_CLI_INVALID_ACCOUNT.(*ext.ClientError).Code, resObj.Code)
}
