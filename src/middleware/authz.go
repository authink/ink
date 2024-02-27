package middleware

import (
	"strconv"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func AuthZ(obj, act string) gin.HandlerFunc {
	return inkstone.HandlerAdapter(func(c *inkstone.Context) {
		var app = c.MustGet("app").(*model.App)

		switch app.Name {
		case env.AppNameAdmin():
			var staff = c.MustGet("account").(*model.Staff)

			sub := strconv.Itoa(int(staff.Id))
			if staff.Super {
				sub = "root"
			}

			dom := strconv.Itoa(int(app.Id))

			if ok, err := authz.RBACEnforcer().Enforce(sub, dom, obj, act); err != nil || !ok {
				c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
			} else {
				c.Next()
			}

		default:
			c.AbortWithForbidden(errors.ERR_UNSUPPORTED_APP)
		}
	})
}
