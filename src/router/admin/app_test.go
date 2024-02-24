package admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
	"github.com/stretchr/testify/assert"
)

func getApps(accessToken string, resObj any) (*httptest.ResponseRecorder, error) {
	return inkstone.TestFetch(
		ctx,
		"GET",
		"admin/apps",
		nil,
		resObj,
		accessToken,
	)
}

func TestApps(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var apps []appRes
	w2, _ := getApps(resObj.AccessToken, &apps)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, len(apps))
	assert.Equal(t, env.AppNameAdmin(), apps[0].Name)
	assert.True(t, apps[0].Active)
}

func tAddApp(accessToken, name string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &addAppReq{name}
	return inkstone.TestFetch(
		ctx,
		"POST",
		"admin/apps",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestAddApp(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resAddApp := &addAppRes{}
	w2, _ := tAddApp(resObj.AccessToken, "appmock", &resAddApp)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Less(t, 100000, resAddApp.Id)
	assert.Equal(t, "appmock", resAddApp.Name)
	assert.NotEmpty(t, resAddApp.Secret)
}
