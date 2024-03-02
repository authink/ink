package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errs"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func setupPermissionGroup(gAdmin *gin.RouterGroup) {
	gPerms := gAdmin.Group(authz.Permissons.Name)
	gPerms.Use(middleware.Authz(authz.Permissons))
	gPerms.GET("", web.HandlerAdapter(permissions))
}

type permissionReq struct {
	AppId int `json:"appId" form:"appId" binding:"required,min=100000" example:"100000"`
}

type permissionRes struct {
	Name     string   `json:"name,omitempty"`
	Resource string   `json:"resource,omitempty"`
	Acts     []string `json:"acts,omitempty"`
	NeedRoot bool     `json:"needRoot"`
}

// permissions godoc
//
//	@Summary		Show permissions
//	@Description	Show permissions
//	@Tags			admin_permissions
//	@Router			/admin/permissions	[get]
//	@Security		ApiKeyAuth
//	@Param			appId	query		int	true	"appId"
//	@Success		200		{array}		permissionRes
//	@Failure		400		{object}	web.ClientError
//	@Failure		401		{object}	web.ClientError
//	@Failure		403		{object}	web.ClientError
//	@Failure		500		{string}	empty
func permissions(c *web.Context) {
	req := new(permissionReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	app, err := orm.App(c.AppContext()).Get(req.AppId)
	if err != nil {
		c.AbortWithServerError(err)
	}

	var res = []permissionRes{}
	for _, v := range authz.ObjList() {
		if v.AppName == app.Name {
			res = append(res, permissionRes{
				Name:     v.Name,
				Resource: v.Resource(),
				Acts:     v.Acts,
				NeedRoot: v.NeedRoot,
			})
		}
	}

	c.Response(res)
}
