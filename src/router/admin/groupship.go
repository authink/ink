package admin

import (
	"strconv"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupGroupshipGroup(gAdmin *gin.RouterGroup) {
	gGroupships := gAdmin.Group(authz.Groupships.Name)
	gGroupships.Use(middleware.Authz(authz.Groupships))
	gGroupships.GET("", inkstone.HandlerAdapter(groupships))
	gGroupships.POST("", inkstone.HandlerAdapter(addGroupship))
	gGroupships.DELETE("", inkstone.HandlerAdapter(deleteGroupship))
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
//	@Failure		400	{object}	inkstone.ClientError
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func groupships(c *inkstone.Context) {
	req := new(groupshipReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
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
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func addGroupship(c *inkstone.Context) {
	req := new(addGroupshipReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
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
//	@Failure		400	{object}	inkstone.ClientError
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func deleteGroupship(c *inkstone.Context) {
	req := new(addGroupshipReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
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
