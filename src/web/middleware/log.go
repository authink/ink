package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
)

func Log(obj authz.Obj) gin.HandlerFunc {
	return web.HandlerAdapter(func(c *web.Context) {
		act := strings.ToUpper(c.Request.Method)

		if !(act == http.MethodPost || act == http.MethodPut || act == http.MethodDelete) {
			c.Next()
			return
		}

		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		c.Next()

		if statusCode := c.Writer.Status(); statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices {
			var body map[string]any
			if reqBody != nil {
				json.Unmarshal(reqBody, &body)
			}

			var app = c.MustGet("app").(*models.App)
			var staff = c.MustGet("account").(*models.Staff)
			var appCtx = c.AppContext()

			go func() {
				orm.Log(appCtx).Insert(models.NewLog(&models.LogDetail{
					AppId:     int(app.Id),
					StaffId:   int(staff.Id),
					Resource:  obj.Resource(),
					Action:    act,
					PathVars:  c.Params,
					QueryVars: c.Request.URL.Query(),
					Body:      body,
				}))
			}()
		}
	})
}
