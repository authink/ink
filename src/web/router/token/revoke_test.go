package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/inkstone/test"
	"github.com/stretchr/testify/assert"
)

func revokeToken(accessToken, refreshToken string) (*httptest.ResponseRecorder, error) {
	reqObj := &refreshReq{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return test.Fetch(
		ctx,
		http.MethodPost,
		"token/revoke",
		reqObj,
		nil,
		"",
	)
}

func TestRevoke(t *testing.T) {
	resObj := &GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	w2, _ := revokeToken(resObj.AccessToken, resObj.RefreshToken)
	assert.Equal(t, http.StatusOK, w2.Code)
}
