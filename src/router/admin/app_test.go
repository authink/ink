package admin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone/test"
	"github.com/stretchr/testify/assert"
)

func getApps(accessToken string, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodGet,
		"admin/apps",
		nil,
		resObj,
		accessToken,
	)
}

func TestApps(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var apps []appRes
	w2, _ := getApps(resObj.AccessToken, &apps)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 2, len(apps))
	assert.Equal(t, env.AppNameAdmin(), apps[0].Name)
	assert.True(t, apps[0].Active)
}

func tAddApp(accessToken, name string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &addAppReq{name}
	return test.Fetch(
		ctx,
		http.MethodPost,
		"admin/apps",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestAddApp(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resAddApp := new(appRes)
	w2, _ := tAddApp(resObj.AccessToken, "appmock", resAddApp)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Less(t, 100001, resAddApp.Id)
	assert.Equal(t, "appmock", resAddApp.Name)
	assert.NotEmpty(t, resAddApp.Secret)
}

func tUpdateApp(accessToken string, id int, reqObj, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodPut,
		fmt.Sprintf("admin/apps/%d", id),
		reqObj,
		resObj,
		accessToken,
	)
}

func TestResetApp(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resetAppReq := &updateAppReq{
		ResetSecret: true,
	}
	resetAppRes := new(appRes)
	w2, _ := tUpdateApp(resObj.AccessToken, 100001, resetAppReq, resetAppRes)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100001, resetAppRes.Id)
	assert.Equal(t, "devtools", resetAppRes.Name)
	assert.NotEqual(t, "123456", resetAppRes.Secret)
}

func TestToggleApp(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resetAppReq := &updateAppReq{
		ActiveToggle: true,
	}
	toggleAppRes := new(appRes)
	w2, _ := tUpdateApp(resObj.AccessToken, 100001, resetAppReq, toggleAppRes)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100001, toggleAppRes.Id)
	assert.Equal(t, "devtools", toggleAppRes.Name)
	assert.False(t, toggleAppRes.Active)
}
