package token

import (
	"os"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/migrate"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func setup(ink *core.Ink) {
	migrate.Schema(ink, "up")
	migrate.Seed(ink)
	r = gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	})
	SetupTokenGroup(r)
}

func teardown(ink *core.Ink) {
	r = nil
	migrate.Schema(ink, "down")
}

func TestMain(m *testing.M) {
	ink := core.NewInk()
	defer ink.Close()

	setup(ink)

	exitCode := m.Run()

	teardown(ink)

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
