package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func setupLogGroup(gAdmin *gin.RouterGroup) {
	gLogs := gAdmin.Group(authz.Logs.Name)
	gLogs.Use(middleware.Authz(authz.Logs))
	gLogs.GET("", web.HandlerAdapter(logs))
}

type logRes struct {
	web.Response
	*models.LogDetail
}

// logs godoc
//
//	@Summary		Show logs
//	@Description	Show logs
//	@Tags			admin_logs
//	@Router			/admin/logs	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		logRes
//	@Failure		400	{object}	web.ClientError
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func logs(c *web.Context) {
	logs, err := orm.Log(c.AppContext()).Find()
	if err != nil {
		c.AbortWithServerError(err)
	}

	var res = []logRes{}
	for _, v := range logs {
		res = append(res, logRes{
			Response: web.Response{
				Id:        int(v.Id),
				CreatedAt: v.CreatedAt,
			},
			LogDetail: v.GetDetail(),
		})
	}

	c.Response(res)
}
