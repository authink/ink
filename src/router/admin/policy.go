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
	gPolicies.POST("", inkstone.HandlerAdapter(addPolicy))
}

type addPolicyReq struct {
	G   int    `json:"g" form:"g" binding:"required,min=100000" example:"100000"`
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

	if _, err := enforcer.AddPermissionForUser(role, req.Obj, req.Act); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
