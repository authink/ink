package middleware

import (
	"strconv"
	"strings"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

const ROOT string = "root"

func Authz(obj authz.Obj) gin.HandlerFunc {
	return inkstone.HandlerAdapter(func(c *inkstone.Context) {
		var app = c.MustGet("app").(*model.App)

		switch app.Name {
		case env.AppNameAdmin():
			var staff = c.MustGet("account").(*model.Staff)

			if obj.NeedRoot {
				if !staff.Super {
					c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
					return
				}

				c.Next()
				return
			}

			act := strings.ToUpper(c.Request.Method)
			if !obj.Support(act) {
				c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
				return
			}

			sub := strconv.Itoa(int(staff.Id))
			if staff.Super {
				sub = ROOT
			}

			dom := strconv.Itoa(int(app.Id))

			if ok, err := authz.RBACEnforcer().Enforce(sub, dom, obj.Name, act); err != nil || !ok {
				c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
			} else {
				c.Next()
			}

		default:
			c.AbortWithForbidden(errors.ERR_UNSUPPORTED_APP)
		}
	})
}
