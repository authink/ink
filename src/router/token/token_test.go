package token

import (
	"testing"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router/common"
	"github.com/authink/ink.go/src/test"
)

func TestMain(m *testing.M) {
	ink := core.NewInk()
	defer ink.Close()

	router, gApi := common.SetupRouter(ink)
	SetupTokenGroup(gApi)

	test.TestMain(ink, router)(m)
}
