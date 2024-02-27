package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/inkstone"
	"github.com/stretchr/testify/assert"
)

func refreshToken(accessToken, refreshToken string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &refreshReq{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return inkstone.TestFetch(
		ctx,
		http.MethodPost,
		"token/refresh",
		reqObj,
		resObj,
		"",
	)
}

func TestRefresh(t *testing.T) {
	resObj := new(GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resObj2 := new(GrantRes)
	w2, _ := refreshToken(resObj.AccessToken, resObj.RefreshToken, resObj2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEmpty(t, resObj2.AccessToken)
	assert.NotEmpty(t, resObj2.RefreshToken)
	assert.Equal(t, "Bearer", resObj2.TokenType)
	assert.Equal(t, 7200, resObj2.ExpiresIn)
}
