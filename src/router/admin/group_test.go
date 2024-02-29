package admin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink.go/src/router/token"
	"github.com/authink/inkstone"
	"github.com/stretchr/testify/assert"
)

func getGroups(accessToken string, gtype, appId int, resObj any) (*httptest.ResponseRecorder, error) {
	return inkstone.TestFetch(
		ctx,
		http.MethodGet,
		fmt.Sprintf("admin/groups?type=%d&appId=%d&limit=1", gtype, appId),
		nil,
		resObj,
		accessToken,
	)
}

func TestGroups(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var res inkstone.PagingResponse[groupRes]
	w2, _ := getGroups(resObj.AccessToken, 1, 100000, &res)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, res.Total)
	assert.Equal(t, 1, len(res.Items))
}
