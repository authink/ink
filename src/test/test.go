package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/migrate"
	"github.com/gin-gonic/gin"
)

type testCtxKey string

var testCtx context.Context
var key = testCtxKey("router")

func setup(ink *core.Ink, router *gin.Engine) {
	migrate.Schema(ink, "up")
	migrate.Seed(ink)
	testCtx = context.WithValue(
		context.Background(),
		key,
		router,
	)
}

func teardown(ink *core.Ink) {
	testCtx = nil
	migrate.Schema(ink, "down")
}

func TestMain(ink *core.Ink, router *gin.Engine) func(*testing.M) {
	return func(m *testing.M) {
		setup(ink, router)

		exitCode := m.Run()

		teardown(ink)

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}

func TestFetch(method, path string, reqObj, resObj any) (w *httptest.ResponseRecorder, err error) {
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
	r := testCtx.Value(key).(*gin.Engine)
	r.ServeHTTP(w, req)

	if w.Body.Len() > 0 {
		err = json.Unmarshal(w.Body.Bytes(), resObj)
	}
	return
}
