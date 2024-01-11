package token

import (
	"os"
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func setup(ink *core.Ink) {
	r = gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	})
	SetupTokenGroup(r)
}

func teardown() {}

func TestMain(m *testing.M) {
	ink := core.NewInk()
	defer ink.Close()

	setup(ink)

	exitCode := m.Run()

	teardown()

	os.Exit(exitCode)
}
