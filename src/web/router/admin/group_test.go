package admin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink/src/web/router/token"
	"github.com/authink/inkstone/test"
	"github.com/authink/inkstone/web"
	"github.com/stretchr/testify/assert"
)

func getGroups(accessToken string, gtype, appId int, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodGet,
		fmt.Sprintf("admin/groups?type=%d&appId=%d&limit=1", gtype, appId),
		nil,
		resObj,
		accessToken,
	)
}

func TestGroups(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var res web.PagingResponse[groupRes]
	w2, _ := getGroups(resObj.AccessToken, 1, 100000, &res)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, res.Total)
	assert.Equal(t, 1, len(res.Items))
}

func tAddGroup(accessToken, name string, gtype, appId int, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &addGroupReq{
		Name: name,
		groupReq: groupReq{
			Type:  gtype,
			AppId: appId,
		},
	}
	return test.Fetch(
		ctx,
		http.MethodPost,
		"admin/groups",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestAddGroup(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resAddGroup := &groupRes{}
	w2, _ := tAddGroup(resObj.AccessToken, "cto", 1, 100000, resAddGroup)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Less(t, 100000, resAddGroup.Id)
	assert.Equal(t, "cto", resAddGroup.Name)
	assert.Equal(t, 1, resAddGroup.Type)
	assert.Equal(t, 100000, resAddGroup.AppId)
}

func tUpdateGroup(accessToken string, id int, reqObj, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodPut,
		fmt.Sprintf("admin/groups/%d", id),
		reqObj,
		resObj,
		accessToken,
	)
}

func TestUpdateGroup(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	updateReq := &updateGroupReq{
		Name:         "cfo",
		ActiveToggle: true,
	}
	updateRes := &groupRes{}
	w2, _ := tUpdateGroup(resObj.AccessToken, 100000, updateReq, updateRes)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100000, updateRes.Id)
	assert.Equal(t, "cfo", updateRes.Name)
	assert.False(t, updateRes.Active)
}
