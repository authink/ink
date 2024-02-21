package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, &dashboardRes{})
}
