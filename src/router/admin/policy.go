package admin

import (
	"strconv"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupPolicyGroup(gAdmin *gin.RouterGroup) {
	gPolicies := gAdmin.Group(authz.Policies.Name)
	gPolicies.Use(middleware.Authz(authz.Policies))
	gPolicies.GET("", inkstone.HandlerAdapter(policies))
	gPolicies.POST("", inkstone.HandlerAdapter(addPolicy))
}

type policyReq struct {
	G   int `json:"g" form:"g" binding:"required,min=100000" example:"100000"`
	Dom int `json:"dom" form:"dom" binding:"required,min=100000" example:"100000"`
}

type policyRes struct {
	Obj string `json:"obj,omitempty"`
	Act string `json:"act,omitempty"`
}

// policies godoc
//
//	@Summary		Show policies
//	@Description	Show policies
//	@Tags			admin_policy
//	@Router			/admin/policies	[get]
//	@Security		ApiKeyAuth
//	@Param			g	query		int	true	"g"
//	@Param			dom	query		int	true	"dom"
//	@Success		200	{array}		policyRes
//	@Failure		400	{object}	inkstone.ClientError
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func policies(c *inkstone.Context) {
	req := new(policyReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	enforcer := authz.RBACEnforcer()
	g := strconv.Itoa(req.G)
	dom := strconv.Itoa(req.Dom)

	var res = []policyRes{}
	permissions, err := enforcer.GetPermissionsForUser(g, dom)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	for _, v := range permissions {
		res = append(res, policyRes{
			Obj: v[2],
			Act: v[3],
		})
	}

	c.Response(res)
}

type addPolicyReq struct {
	policyReq
	Obj string `json:"obj" form:"obj" binding:"required,min=2" example:"admin.dev/apps"`
	Act string `json:"act" form:"act" binding:"required,eq=GET|eq=POST|eq=PUT|eq=DELETE" example:"GET"`
}

// addPolicy godoc
//
//	@Summary		Add a policy
//	@Description	Add a policy
//	@Tags			admin_policy
//	@Router			/admin/policies	[post]
//	@Security		ApiKeyAuth
//	@Param			addPolicyReq	body		addPolicyReq	true	"request body"
//	@Success		200				{string}	empty
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func addPolicy(c *inkstone.Context) {
	req := new(addPolicyReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	enforcer := authz.RBACEnforcer()
	role := strconv.Itoa(req.G)
	dom := strconv.Itoa(req.Dom)

	if _, err := enforcer.AddPermissionForUser(role, dom, req.Obj, req.Act); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
