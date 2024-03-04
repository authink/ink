package admin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/web/router/token"
	"github.com/authink/inkstone/test"
	"github.com/authink/inkstone/web"
	"github.com/stretchr/testify/assert"
)

func getTokens(accessToken string, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodGet,
		"admin/tokens?limit=2",
		nil,
		resObj,
		accessToken,
	)
}

func TestTokens(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resObj2 := new(token.GrantRes)
	w2, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj2)

	assert.Equal(t, http.StatusOK, w2.Code)
	assert.NotEmpty(t, resObj2.AccessToken)
	assert.NotEmpty(t, resObj2.RefreshToken)

	var res web.PagingResponse[tokenRes]
	w3, _ := getTokens(resObj.AccessToken, &res)
	assert.Equal(t, http.StatusOK, w3.Code)
	assert.Equal(t, 2, len(res.Items))
}

func tDeleteToken(accessToken string, id int) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("admin/tokens/%d", id),
		nil,
		nil,
		accessToken,
	)
}

func TestDeleteToken(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	w2, _ := tDeleteToken(resObj.AccessToken, 100000)
	assert.Equal(t, http.StatusOK, w2.Code)
}
