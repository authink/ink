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
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	var res inkstone.PagingResponse[staffRes]
	w2, _ := getStaffs(resObj.AccessToken, &res)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 1, len(res.Items))
}

func tAddStaff(accessToken, email, phone string, super bool, resObj any) (*httptest.ResponseRecorder, error) {
	reqObj := &addStaffReq{email, phone, super}
	return inkstone.TestFetch(
		ctx,
		"POST",
		"admin/staffs",
		reqObj,
		resObj,
		accessToken,
	)
}

func TestAddStaff(t *testing.T) {
	resObj := new(token.GrantRes)
	w, _ := grantToken(100000, "123456", "admin@huoyijie.cn", "123456", resObj)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)

	resAddStaff := new(addStaffRes)
	w2, _ := tAddStaff(resObj.AccessToken, "example@huoyijie.cn", "18555201314", false, resAddStaff)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, 100001, resAddStaff.Id)
	assert.Equal(t, "example@huoyijie.cn", resAddStaff.Email)
	assert.NotEmpty(t, resAddStaff.Password)
}
