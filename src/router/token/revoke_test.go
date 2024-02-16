package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/ext"
	"github.com/stretchr/testify/assert"
)

func revokeToken(accessToken, refreshToken string) (*httptest.ResponseRecorder, error) {
	reqObj := &refreshReq{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ext.TestFetch(
		r,
		"POST",
		"/api/v1/token/revoke",
		reqObj,
		nil,
	)
}

func TestRevoke(t *testing.T) {
	resObj := &grantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	w2, _ := revokeToken(resObj.AccessToken, resObj.RefreshToken)
	assert.Equal(t, http.StatusOK, w2.Code)
}
