package admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/authink/ink/src/web/router/token"
	"github.com/authink/stone/test"
	"github.com/authink/stone/web"
	"github.com/stretchr/testify/assert"
)

func getStaffs(accessToken string, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodGet,
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

	var res web.PagingResponse[staffRes]
	w2, _ := getStaffs(resObj.AccessToken, &res)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, len(res.Items))
}

func tAddStaff(accessToken, email, phone string, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &addStaffReq{email, phone}
	return test.Fetch(
		ctx,
		http.MethodPost,
		"admin/staffs",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestAddStaff(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resAddStaff := &staffRes{}
	w2, _ := tAddStaff(resObj.AccessToken, "example@huoyijie.cn", "18555201314", resAddStaff)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100002, resAddStaff.Id)
	assert.Equal(t, "example@huoyijie.cn", resAddStaff.Email)
	assert.NotEmpty(t, resAddStaff.Password)
}

func tUpdateStaff(accessToken string, reqObj, resObj any) (*httptest.ResponseRecorder, error) {
	return test.Fetch(
		ctx,
		http.MethodPut,
		"admin/staffs",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestUpdateStaff(t *testing.T) {
	resObj := &token.GrantRes{}
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	updateReq := &updateStaffReq{
		Id:              100001,
		Phone:           "12112112112",
		ActiveToggle:    true,
		DepartureToggle: true,
		ResetPassword:   true,
	}
	updateRes := &staffRes{}
	w2, _ := tUpdateStaff(resObj.AccessToken, updateReq, updateRes)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100001, updateRes.Id)
	assert.Equal(t, "12112112112", updateRes.Phone)
	assert.NotEqual(t, "123456", updateRes.Password)
	assert.False(t, updateRes.Active)
	assert.False(t, updateRes.Super)
	assert.True(t, updateRes.Departure)
}
