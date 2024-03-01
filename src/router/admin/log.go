package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupLogGroup(gAdmin *gin.RouterGroup) {
	gLogs := gAdmin.Group(authz.Logs.Name)
	gLogs.Use(middleware.Authz(authz.Logs))
	gLogs.GET("", inkstone.HandlerAdapter(logs))
}

type logRes struct {
	inkstone.Response
	*model.LogDetail
}

// logs godoc
//
//	@Summary		Show logs
//	@Description	Show logs
//	@Tags			admin_logs
//	@Router			/admin/logs	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		logRes
//	@Failure		400	{object}	inkstone.ClientError
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func logs(c *inkstone.Context) {
	logs, err := orm.Log(c.AppContext()).Find()
	if err != nil {
		c.AbortWithServerError(err)
	}

	var res = []logRes{}
	for _, v := range logs {
		res = append(res, logRes{
			Response: inkstone.Response{
				Id:        int(v.Id),
				CreatedAt: v.CreatedAt,
			},
			LogDetail: v.GetDetail(),
		})
	}

	c.Response(res)
}
