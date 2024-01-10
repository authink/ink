package ext

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/authink/ink.go/src/core"
	"github.com/gin-gonic/gin"
)

type SetupGroup func(*gin.Engine)

func TestFetch(method, path string, reqObj, resObj any, setupGroup SetupGroup) (w *httptest.ResponseRecorder, err error) {
	ink := core.NewInk()
	defer ink.Close()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	})
	setupGroup(r)

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

	if w.Body.Len() > 0 {
		err = json.Unmarshal(w.Body.Bytes(), resObj)
	}
	return
}
