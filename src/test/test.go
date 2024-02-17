package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/migrate"
	"github.com/gin-gonic/gin"
)

type ContextKey string

var InkKey = ContextKey("ink")
var RouterKey = ContextKey("router")

func setup(ink *core.Ink) {
	migrate.Schema(ink, "up")
	migrate.Seed(ink)
}

func teardown(ink *core.Ink) {
	migrate.Schema(ink, "down")
}

func Main(ctx *context.Context, ink *core.Ink, router *gin.Engine) func(*testing.M) {
	*ctx = context.WithValue(
		*ctx,
		InkKey,
		ink,
	)
	*ctx = context.WithValue(
		*ctx,
		RouterKey,
		router,
	)

	return func(m *testing.M) {
		setup(ink)

		exitCode := m.Run()

		teardown(ink)

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}

func Fetch(ctx context.Context, method, pathname string, reqObj, resObj any, accessToken string) (w *httptest.ResponseRecorder, err error) {
	ink := ctx.Value(InkKey).(*core.Ink)

	var reader io.Reader
	if reqObj != nil {
		reqBody, _ := json.Marshal(reqObj)
		reader = strings.NewReader(string(reqBody))
	}

	w = httptest.NewRecorder()
	req, _ := http.NewRequest(
		method,
		path.Join("/", ink.Env.BasePath, pathname),
		reader,
	)

	if accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	router := ctx.Value(RouterKey).(*gin.Engine)
	router.ServeHTTP(w, req)

	if w.Body.Len() > 0 {
		err = json.Unmarshal(w.Body.Bytes(), resObj)
	}
	return
}
