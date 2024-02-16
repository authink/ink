package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// dashboard godoc
//
//	@Summary		Show dashboard
//	@Description	Show dashboard
//	@Tags			dashboard
//	@Produce		json
//	@Success		200	{object}	map[string]any
//	@Failure		400	{object}	map[string]any
//	@Failure		404	{object}	map[string]any
//	@Failure		500	{object}	map[string]any
//	@Router			/admin/dashboard [get]
//
//	@Security		ApiKeyAuth
func dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
