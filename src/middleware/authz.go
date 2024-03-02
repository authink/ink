package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

const ROOT string = "root"

func Authz(obj authz.Obj) gin.HandlerFunc {
	return web.HandlerAdapter(func(c *web.Context) {
		var app = c.MustGet("app").(*model.App)

		switch app.Name {
		case env.AppNameAdmin():
			var staff = c.MustGet("account").(*model.Staff)
			act := strings.ToUpper(c.Request.Method)

			if obj.NeedRoot {
				if !staff.Super {
					c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
					return
				}
			} else {
				if !obj.Support(act) {
					c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
					return
				}

				sub := strconv.Itoa(int(staff.Id))
				if staff.Super {
					sub = ROOT
				}

				dom := strconv.Itoa(int(app.Id))

				if ok, err := authz.RBACEnforcer().Enforce(sub, dom, obj.Resource(), act); err != nil || !ok {
					c.AbortWithForbidden(errors.ERR_NO_PERMISSION)
					return
				}
			}

			c.Next()
			if statusCode := c.Writer.Status(); statusCode == http.StatusOK {
				orm.Log(c.AppContext()).Insert(model.NewLog(&model.LogDetail{
					AppId:     int(app.Id),
					StaffId:   int(staff.Id),
					Resource:  obj.Resource(),
					Action:    act,
					PathVars:  c.Params,
					QueryVars: c.Request.URL.Query(),
				}))
			}

		default:
			c.AbortWithForbidden(errors.ERR_UNSUPPORTED_APP)
		}
	})
}
