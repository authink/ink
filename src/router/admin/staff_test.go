package admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
	"github.com/stretchr/testify/assert"
)

func getStaffs(accessToken string, resObj any) (*httptest.ResponseRecorder, error) {
	return inkstone.TestFetch(
		ctx,
		"GET",
		"admin/staffs?limit=1",
		nil,
		resObj,
		accessToken,
	)
}

func TestStaffs(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var res pageRes[staffRes]
	w2, _ := getStaffs(resObj.AccessToken, &res)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, len(res.Items))
}
