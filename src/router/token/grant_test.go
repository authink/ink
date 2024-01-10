package token

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest(ink *core.Ink) (r *gin.Engine) {
	r = gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	})

	SetupTokenGroup(r)
	return
}

func request(method, path string, reqObj, resObj any) (w *httptest.ResponseRecorder, err error) {
	ink := core.NewInk()
	defer ink.Close()
	r := setupTest(ink)

	var reader io.Reader
	if reqObj != nil {
		reqBody, _ := json.Marshal(reqObj)
		reader = strings.NewReader(string(reqBody))
	}

	w = httptest.NewRecorder()
	req, _ := http.NewRequest(
		method,
		path,
		reader,
	)
	r.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), resObj)
	return
}

func TestTokenGrant(t *testing.T) {
	reqObj := &reqGrant{
		AppId:     100000,
		AppSecret: "123456",
		Email:     "admin@huoyijie.cn",
		Password:  "123456",
	}

	resObj := &resGrant{}

	w, err := request(
		"POST",
		"/token/grant",
		reqObj,
		resObj,
	)
	if err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Bearer", resObj.TokenType)
	assert.Equal(t, 7200, resObj.ExpiresIn)
	assert.NotEmpty(t, resObj.AccessToken)
	assert.NotEmpty(t, resObj.RefreshToken)
}
