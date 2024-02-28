package admin

import (
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupDashboard(gAdmin *gin.RouterGroup) {
	gAdmin.GET("dashboard", inkstone.HandlerAdapter(dashboard))
}

type dashboardRes struct {
	Count int `json:"count"`
}

// dashboard godoc
//
//	@Summary		Show dashboard
//	@Description	Show dashboard
//	@Tags			dashboard
//	@Router			/admin/dashboard [get]
//	@Security		ApiKeyAuth
//	@Param			category	query		string	true	"staff"	Enums(staff, user)
//	@Success		200			{object}	dashboardRes
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func dashboard(c *inkstone.Context) {
	c.Response(new(dashboardRes))
}
