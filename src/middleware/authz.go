package middleware

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/inkstone"
)

func AuthZ() inkstone.HandlerFunc {
	return func(c *inkstone.Context) {
		c.Set("enforcer", authz.RBACEnforcer())
		c.Next()
	}
}
