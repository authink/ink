package admin

import (
	"strconv"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/web/errs"
	"github.com/authink/ink.go/src/web/middleware"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func setupGroupshipGroup(gAdmin *gin.RouterGroup) {
	gGroupships := gAdmin.Group(authz.Groupships.Name)
	gGroupships.Use(middleware.Authz(authz.Groupships))
	gGroupships.GET("", web.HandlerAdapter(groupships))
	gGroupships.POST("", web.HandlerAdapter(addGroupship))
	gGroupships.DELETE("", web.HandlerAdapter(deleteGroupship))
}

type groupshipReq struct {
	Sub int `json:"sub" form:"sub" binding:"required,min=100000" example:"100000"`
	Dom int `json:"dom" form:"dom" binding:"required,min=100000" example:"100000"`
}

// groupships godoc
//
//	@Summary		Show groupships
//	@Description	Show groupships
//	@Tags			admin_groupship
//	@Router			/admin/groupships	[get]
//	@Security		ApiKeyAuth
//	@Param			sub	query		int	true	"sub"
//	@Param			dom	query		int	true	"dom"
//	@Success		200	{array}		string
//	@Failure		400	{object}	web.ClientError
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func groupships(c *web.Context) {
	req := &groupshipReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	enforcer := authz.RBACEnforcer()
	sub := strconv.Itoa(req.Sub)
	dom := strconv.Itoa(req.Dom)

	var res = []string{}
	if roles := enforcer.GetRolesForUserInDomain(sub, dom); len(roles) > 0 {
		res = roles
	}

	c.Response(res)
}

type addGroupshipReq struct {
	groupshipReq
	G int `json:"g" form:"g" binding:"required,min=100000" example:"100000"`
}

// addGroupship godoc
//
//	@Summary		Add a groupship
//	@Description	Add a groupship
//	@Tags			admin_groupship
//	@Router			/admin/groupships	[post]
//	@Security		ApiKeyAuth
//	@Param			addGroupshipReq	body		addGroupshipReq	true	"request body"
//	@Success		200				{string}	empty
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func addGroupship(c *web.Context) {
	req := &addGroupshipReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	enforcer := authz.RBACEnforcer()
	sub := strconv.Itoa(req.Sub)
	dom := strconv.Itoa(req.Dom)
	role := strconv.Itoa(req.G)
	if _, err := enforcer.AddRoleForUserInDomain(sub, role, dom); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}

// deleteGroupship godoc
//
//	@Summary		Delete a groupship
//	@Description	Delete a groupship
//	@Tags			admin_groupship
//	@Router			/admin/groupships	[delete]
//	@Security		ApiKeyAuth
//	@Param			sub	query		int	true	"sub"
//	@Param			dom	query		int	true	"dom"
//	@Param			g	query		int	true	"g"
//	@Success		200	{string}	empty
//	@Failure		400	{object}	web.ClientError
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func deleteGroupship(c *web.Context) {
	req := &addGroupshipReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	enforcer := authz.RBACEnforcer()
	sub := strconv.Itoa(req.Sub)
	dom := strconv.Itoa(req.Dom)
	role := strconv.Itoa(req.G)

	if _, err := enforcer.DeleteRoleForUserInDomain(sub, role, dom); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
