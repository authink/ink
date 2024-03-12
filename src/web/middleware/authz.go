package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
)

const ROOT string = "root"

func Authz(obj authz.Obj) gin.HandlerFunc {
	return web.HandlerAdapter(func(c *web.Context) {
		var app = c.MustGet("app").(*models.App)

		switch app.Name {
		case envs.AppNameAdmin():
			var staff = c.MustGet("account").(*models.Staff)
			act := strings.ToUpper(c.Request.Method)

			if obj.NeedRoot {
				if !staff.Super {
					c.AbortWithForbidden(errs.ERR_NO_PERMISSION)
					return
				}
			} else {
				if !obj.Support(act) {
					c.AbortWithForbidden(errs.ERR_NO_PERMISSION)
					return
				}

				sub := strconv.Itoa(int(staff.Id))
				if staff.Super {
					sub = ROOT
				}

				dom := strconv.Itoa(int(app.Id))

				if ok, err := authz.RBACEnforcer().Enforce(sub, dom, obj.Resource(), act); err != nil || !ok {
					c.AbortWithForbidden(errs.ERR_NO_PERMISSION)
					return
				}
			}

			c.Next()
			if statusCode := c.Writer.Status(); statusCode == http.StatusOK {
				orm.Log(c.AppContext()).Insert(models.NewLog(&models.LogDetail{
					AppId:     int(app.Id),
					StaffId:   int(staff.Id),
					Resource:  obj.Resource(),
					Action:    act,
					PathVars:  c.Params,
					QueryVars: c.Request.URL.Query(),
				}))
			}

		default:
			c.AbortWithForbidden(errs.ERR_UNSUPPORTED_APP)
		}
	})
}
