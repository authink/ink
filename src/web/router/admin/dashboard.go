package admin

import (
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
)

func setupDashboard(gAdmin *gin.RouterGroup) {
	gAdmin.GET("dashboard", web.HandlerAdapter(dashboard))
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
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func dashboard(c *web.Context) {
	c.Response(&dashboardRes{})
}
